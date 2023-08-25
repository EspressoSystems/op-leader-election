package derive

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"

	"github.com/ethereum-optimism/optimism/op-bindings/bindings"
	"github.com/ethereum/go-ethereum/accounts/abi"
)

// Frames cannot be larger than 1 MB.
// Data transactions that carry frames are generally not larger than 128 KB due to L1 network conditions,
// but we leave space to grow larger anyway (gas limit allows for more data).
const MaxFrameLen = 1_000_000

var SubmitAbi = loadSubmitAbi()

func loadSubmitAbi() abi.Method {
	abi, err := bindings.LeaderElectionBatchInboxMetaData.GetAbi()
	if err != nil || abi == nil {
		log.Fatalf("could not get LeaderElectionBatchInbox contract abi")
	}
	submit, exists := abi.Methods["submit"]
	if !exists {
		log.Fatalf("no submit method in abi")
	}
	return submit
}

// Data Format
//
// frame = channel_id ++ frame_number ++ frame_data_length ++ frame_data ++ is_last
//
// channel_id        = bytes16
// frame_number      = uint16
// frame_data_length = uint32
// frame_data        = bytes
// is_last           = bool

type Frame struct {
	ID          ChannelID `json:"id"`
	FrameNumber uint16    `json:"frame_number"`
	Data        []byte    `json:"data"`
	IsLast      bool      `json:"is_last"`
}

// MarshalBinary writes the frame to `w`.
// It returns any errors encountered while writing, but
// generally expects the writer very rarely fail.
func (f *Frame) MarshalBinary(w io.Writer) error {
	_, err := w.Write(f.ID[:])
	if err != nil {
		return err
	}
	if err := binary.Write(w, binary.BigEndian, f.FrameNumber); err != nil {
		return err
	}
	if err := binary.Write(w, binary.BigEndian, uint32(len(f.Data))); err != nil {
		return err
	}
	_, err = w.Write(f.Data)
	if err != nil {
		return err
	}
	if f.IsLast {
		if _, err = w.Write([]byte{1}); err != nil {
			return err
		}
	} else {
		if _, err = w.Write([]byte{0}); err != nil {
			return err
		}
	}
	return nil
}

type ByteReader interface {
	io.Reader
	io.ByteReader
}

// UnmarshalBinary consumes a full frame from the reader.
// If `r` fails a read, it returns the error from the reader
// The reader will be left in a partially read state.
//
// If r doesn't return any bytes, returns io.EOF.
// If r unexpectedly stops returning data half-way, returns io.ErrUnexpectedEOF.
func (f *Frame) UnmarshalBinary(r ByteReader) error {
	if _, err := io.ReadFull(r, f.ID[:]); err != nil {
		// Forward io.EOF here ok, would mean not a single byte from r.
		return fmt.Errorf("reading channel_id: %w", err)
	}
	if err := binary.Read(r, binary.BigEndian, &f.FrameNumber); err != nil {
		return fmt.Errorf("reading frame_number: %w", eofAsUnexpectedMissing(err))
	}

	var frameLength uint32
	if err := binary.Read(r, binary.BigEndian, &frameLength); err != nil {
		return fmt.Errorf("reading frame_data_length: %w", eofAsUnexpectedMissing(err))
	}

	// Cap frame length to MaxFrameLen (currently 1MB)
	if frameLength > MaxFrameLen {
		return fmt.Errorf("frame_data_length is too large: %d", frameLength)
	}
	f.Data = make([]byte, int(frameLength))
	if _, err := io.ReadFull(r, f.Data); err != nil {
		return fmt.Errorf("reading frame_data: %w", eofAsUnexpectedMissing(err))
	}

	if isLastByte, err := r.ReadByte(); err != nil {
		return fmt.Errorf("reading final byte (is_last): %w", eofAsUnexpectedMissing(err))
	} else if isLastByte == 0 {
		f.IsLast = false
	} else if isLastByte == 1 {
		f.IsLast = true
	} else {
		return errors.New("invalid byte as is_last")
	}
	return nil
}

// eofAsUnexpectedMissing converts an io.EOF in the error chain of err into an
// io.ErrUnexpectedEOF. It should be used to convert intermediate io.EOF errors
// in unmarshaling code to achieve idiomatic error behavior.
// Other errors are passed through unchanged.
func eofAsUnexpectedMissing(err error) error {
	if errors.Is(err, io.EOF) {
		return fmt.Errorf("fully missing: %w", io.ErrUnexpectedEOF)
	}
	return err
}

// Frames are stored in L1 transactions with the following format:
// data = DerivationVersion0 ++ Frame(s)
// Where there is one or more frames concatenated together.

// ParseFrames parse the on chain serialization of frame(s) in
// an L1 transaction. Currently only version 0 of the serialization
// format is supported.
// All frames must be parsed without error and there must not be
// any left over data and there must be at least one frame.
func ParseFrames(data []byte) ([]Frame, error) {
	if len(data) == 0 {
		return nil, errors.New("data array must not be empty")
	}
	// Check if the calldata signature matches the submit function to determine
	// if this was a batch that was sent to the inbox smart contract. If that's
	// the case, we need to decode the arguments and, later check if the meta
	// data matches the decoded frame. Otherwise, treat it as a frame that was
	// sent to the (non-contract) batch inbox address.
	var metas []bindings.LeaderElectionBatchInboxMeta
	isFromContract := len(data) >= 4 && bytes.Equal(data[:4], SubmitAbi.ID)
	if isFromContract {
		var err error
		metas, data, err = ParseFramesV2(data)
		if err != nil {
			return nil, fmt.Errorf("parsing v2 frames: %w", err)
		}
	}

	if data[0] != DerivationVersion0 {
		return nil, fmt.Errorf("invalid derivation format byte: got %d", data[0])
	}
	buf := bytes.NewBuffer(data[1:])
	var frames []Frame
	for buf.Len() > 0 {
		var f Frame
		if err := f.UnmarshalBinary(buf); err != nil {
			return nil, fmt.Errorf("parsing frame %d: %w", len(frames), err)
		}
		frames = append(frames, f)
	}
	if buf.Len() != 0 {
		return nil, fmt.Errorf("did not fully consume data: have %d frames and %d bytes left", len(frames), buf.Len())
	}
	if len(frames) == 0 {
		return nil, errors.New("was not able to find any frames")
	}

	// Check that the metadata matches the frames
	if isFromContract {
		if len(frames) != len(metas) {
			return nil, fmt.Errorf("number of frames and metas do not match: %d != %d", len(frames), len(metas))
		}
		for i := range frames {
			frame := frames[i]
			meta := metas[i]
			if frame.ID != meta.ChannelId {
				return nil, fmt.Errorf("frame %d channel id does not match meta: %s != %s", i, frame.ID, meta.ChannelId)
			}
			if frame.FrameNumber != meta.FrameNumber {
				return nil, fmt.Errorf("frame %d frame number does not match meta: %d != %d", i, frame.FrameNumber, meta.FrameNumber)
			}
			if len(frame.Data) != int(meta.FrameDataLength) {
				return nil, fmt.Errorf("frame %d frame data length does not match meta: %d != %d", i, len(frame.Data), meta.FrameDataLength)
			}
			if frame.IsLast != meta.IsLast {
				return nil, fmt.Errorf("frame %d isLast does not match meta: %t != %t", i, frame.IsLast, meta.IsLast)
			}
		}
	}

	return frames, nil
}

// TODO currently untested
func ParseFramesV2(data []byte) ([]bindings.LeaderElectionBatchInboxMeta, []byte, error) {

	decoded, err := SubmitAbi.Inputs.Unpack(data)
	if err != nil {
		return nil, nil, fmt.Errorf("could not decode data: %w", err)
	}
	metas, ok := decoded[0].([]bindings.LeaderElectionBatchInboxMeta)
	if !ok {
		return nil, nil, fmt.Errorf("could not decode metas")
	}
	frames, ok := decoded[1].([]byte)
	if !ok {
		return nil, nil, fmt.Errorf("could not decode frames")
	}

	return metas, frames, nil
}
