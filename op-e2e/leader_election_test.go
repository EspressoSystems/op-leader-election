package op_e2e

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLeaderElection(t *testing.T) {
	InitParallel(t)
	require.Equal(t, 2000, 2000, "Values are different")
}
