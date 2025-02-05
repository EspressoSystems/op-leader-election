package derive

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func TestBatchRoundTrip(t *testing.T) {
	batches := []*BatchData{
		{
			Version: BatchV1Type,
			BatchV2: BatchV2{common.Address{}, BatchV1{
				ParentHash:   common.Hash{},
				EpochNum:     0,
				Timestamp:    0,
				Transactions: []hexutil.Bytes{},
			}},
		},
		{
			Version: BatchV1Type,
			BatchV2: BatchV2{common.Address{}, BatchV1{
				ParentHash:   common.Hash{31: 0x42},
				EpochNum:     1,
				Timestamp:    1647026951,
				Transactions: []hexutil.Bytes{[]byte{0, 0, 0}, []byte{0x76, 0xfd, 0x7c}},
			}},
		},
	}

	for i, batch := range batches {
		enc, err := batch.MarshalBinary()
		assert.NoError(t, err)
		var dec BatchData
		err = dec.UnmarshalBinary(enc)
		assert.NoError(t, err)
		assert.Equal(t, batch, &dec, "Batch not equal test case %v", i)
	}
}
