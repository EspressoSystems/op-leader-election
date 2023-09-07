package op_e2e

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum-optimism/optimism/op-bindings/bindings"
	"github.com/ethereum-optimism/optimism/op-node/testlog"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/log"
	"github.com/stretchr/testify/require"
)

func TestLeaderElection(t *testing.T) {
	InitParallel(t)
	
	cfg := DefaultSystemConfig(t)

	sys, err := cfg.Start(t)
	require.Nil(t, err, "Error starting up system")
	defer sys.Close()

	log := testlog.Logger(t, log.LvlInfo)
	log.Info("genesis", "l2", sys.RollupConfig.Genesis.L2, "l1", sys.RollupConfig.Genesis.L1, "l2_time", sys.RollupConfig.Genesis.L2Time)

	l1Client := sys.Clients["l1"]

	opts, err := bind.NewKeyedTransactorWithChainID(sys.cfg.Secrets.Alice, cfg.L1ChainIDBig())
	log.Info(opts.GasPrice.String())
	require.Nil(t, err)

	// Check that interacting with the contract fails with some hardcoded address
	leaderElectionContractWrongAddress := common.BytesToAddress([]byte("0x"))
	leaderElectionContractWrong, err := bindings.NewLeaderElectionBatchInbox(leaderElectionContractWrongAddress, l1Client)
	require.Nil(t, err)
	aliceAddress := sys.cfg.Secrets.Addresses().Alice
	blockNumber := big.NewInt(0)
	_, err = leaderElectionContractWrong.IsCurrentLeader(&bind.CallOpts{}, aliceAddress, blockNumber)
	require.Error(t, err)

	// Now with the address from sys.cfg we can interact with the contract
	leaderElectionContractAddress := sys.cfg.L1Deployments.RoundRobinLeaderElection
	log.Info("leaderElectionContractAddress: %s", leaderElectionContractAddress.String())
	leaderElectionContractNew, err := bindings.NewLeaderElectionBatchInbox(leaderElectionContractAddress, l1Client)
	require.Nil(t, err)

	isLeader, err := leaderElectionContractNew.IsCurrentLeader(&bind.CallOpts{}, aliceAddress, blockNumber)
	require.Nil(t, err)
	require.False(t, isLeader)

}
