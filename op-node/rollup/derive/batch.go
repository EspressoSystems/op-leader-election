package derive

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"sync"

	"github.com/ethereum-optimism/optimism/op-node/rollup"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rlp"
)

// Batch format
// first byte is type followed by bytestring.
//
// BatchV1Type := 0
// batchV1 := BatchV1Type ++ RLP([epoch, timestamp, transaction_list])
//
// An empty input is not a valid batch.
//
// Note: the type system is based on L1 typed transactions.
//
// BatchV2Type := 1
// batchV2 := BatchV2Type ++ RLP([fee_addr, epoch, timestamp, transaction_list])

// encodeBufferPool holds temporary encoder buffers for batch encoding
var encodeBufferPool = sync.Pool{
	New: func() any { return new(bytes.Buffer) },
}

const (
	BatchV1Type = iota
	BatchV2Type = BatchV1Type + 1
)

type BatchV1 struct {
	ParentHash common.Hash  // parent L2 block hash
	EpochNum   rollup.Epoch // aka l1 num
	EpochHash  common.Hash  // block hash
	Timestamp  uint64
	// no feeRecipient address input, all fees go to a L2 contract
	Transactions []hexutil.Bytes
}

type BatchV2 struct {
	PayToAddr common.Address
	BatchV1
}

type BatchData struct {
	Version int
	BatchV2
	// batches may contain additional data with new upgrades
}

func (b *BatchV1) Epoch() eth.BlockID {
	return eth.BlockID{Hash: b.EpochHash, Number: uint64(b.EpochNum)}
}

func (b *BatchV2) Epoch() eth.BlockID {
	return eth.BlockID{Hash: b.EpochHash, Number: uint64(b.EpochNum)}
}

// EncodeRLP implements rlp.Encoder
func (b *BatchData) EncodeRLP(w io.Writer) error {
	buf := encodeBufferPool.Get().(*bytes.Buffer)
	defer encodeBufferPool.Put(buf)
	buf.Reset()
	if err := b.encodeTyped(buf); err != nil {
		return err
	}
	return rlp.Encode(w, buf.Bytes())
}

// MarshalBinary returns the canonical encoding of the batch.
func (b *BatchData) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	err := b.encodeTyped(&buf)
	return buf.Bytes(), err
}

func (b *BatchData) encodeTyped(buf *bytes.Buffer) error {
	if b.Version == BatchV1Type {
		buf.WriteByte(BatchV1Type)
		return rlp.Encode(buf, &b.BatchV1)
	} else if b.Version == BatchV2Type {
		buf.WriteByte(BatchV2Type)
		return rlp.Encode(buf, &b.BatchV2)
	} else {
		return fmt.Errorf("unrecognized batch type: %d", b.Version)
	}
}

// DecodeRLP implements rlp.Decoder
func (b *BatchData) DecodeRLP(s *rlp.Stream) error {
	if b == nil {
		return errors.New("cannot decode into nil BatchData")
	}
	v, err := s.Bytes()
	if err != nil {
		return err
	}
	return b.decodeTyped(v)
}

// UnmarshalBinary decodes the canonical encoding of batch.
func (b *BatchData) UnmarshalBinary(data []byte) error {
	if b == nil {
		return errors.New("cannot decode into nil BatchData")
	}
	return b.decodeTyped(data)
}

func (b *BatchData) decodeTyped(data []byte) error {
	if len(data) == 0 {
		return fmt.Errorf("batch too short")
	}
	switch data[0] {
	case BatchV1Type:
		b.Version = BatchV1Type
		b.BatchV2.PayToAddr = common.Address{}
		return rlp.DecodeBytes(data[1:], &b.BatchV1)
	case BatchV2Type:
		b.Version = BatchV2Type
		return rlp.DecodeBytes(data[1:], &b.BatchV2)
	default:
		return fmt.Errorf("unrecognized batch type: %d", data[0])
	}
}
