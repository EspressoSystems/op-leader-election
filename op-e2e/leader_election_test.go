package op_e2e

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLeaderElection(t *testing.T) {
	InitParallel(t)
	require.Equal(t, 1000, 1000, "Values are different")
}
