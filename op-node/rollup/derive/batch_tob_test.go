package derive

import (
	"testing"

	"github.com/ethereum-optimism/optimism/op-node/testutils/fuzzerutils"
	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/require"
)

// FuzzBatchRoundTrip executes a fuzz test similar to TestBatchRoundTrip, which tests that arbitrary BatchData will be
// encoded and decoded without loss of its original values.
func FuzzBatchRoundTrip(f *testing.F) {
	f.Fuzz(func(t *testing.T, fuzzedData []byte) {
		// Create our fuzzer wrapper to generate complex values
		typeProvider := fuzz.NewFromGoFuzz(fuzzedData).NilChance(0).MaxDepth(10000).NumElements(0, 0x100).AllowUnexportedFields(true)
		fuzzerutils.AddFuzzerFunctions(typeProvider)

		// Add an extra fuzzer that respects constraints of V1 vs V2 batch
		typeProvider.Funcs(
			func(e *BatchData, c fuzz.Continue) {
				// Only 0 and 1 are valid versions.
				e.Version = c.Intn(2)

				// Fuzz the PayToAddr for v2 batches
				if e.Version == 1 {
					c.Fuzz(&e.BatchV2.PayToAddr)
				}

				// Fuzz the rest of the batch data
				c.Fuzz(&e.BatchV2.BatchV1)
			},
		)

		// Create our batch data from fuzzed data
		var batchData BatchData
		typeProvider.Fuzz(&batchData)

		// Encode our batch data
		enc, err := batchData.MarshalBinary()
		require.NoError(t, err)

		// Decode our encoded batch data
		var dec BatchData
		err = dec.UnmarshalBinary(enc)
		require.NoError(t, err)

		// Ensure the round trip encoding of batch data did not result in data loss
		require.Equal(t, &batchData, &dec, "round trip batch encoding/decoding did not match original values")
	})
}
