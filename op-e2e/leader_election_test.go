package op_e2e

import (
	"math/big"
	"testing"
	"time"

	"github.com/ethereum-optimism/optimism/op-bindings/bindings"
	"github.com/ethereum-optimism/optimism/op-node/testlog"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/log"
	"github.com/stretchr/testify/require"
)

func TestLeaderElection(t *testing.T) {
	InitParallel(t)
	require.Equal(t, 2000, 2000, "Values are different")

	cfg := DefaultSystemConfig(t)

	sys, err := cfg.Start(t)
	require.Nil(t, err, "Error starting up system")
	defer sys.Close()

	log := testlog.Logger(t, log.LvlInfo)
	log.Info("genesis", "l2", sys.RollupConfig.Genesis.L2, "l1", sys.RollupConfig.Genesis.L1, "l2_time", sys.RollupConfig.Genesis.L2Time)

	l1Client := sys.Clients["l1"]

	opts, err := bind.NewKeyedTransactorWithChainID(sys.cfg.Secrets.Alice, cfg.L1ChainIDBig())
	require.Nil(t, err)

	leaderElectionContractAddress, tx, LeaderElectionContract, err := bindings.DeployRoundRobinLeaderElection(opts, l1Client)
	require.NoError(t, err)
	_, err = waitForTransaction(tx.Hash(), l1Client, 3*time.Duration(cfg.DeployConfig.L1BlockTime)*time.Second)
	require.NoError(t, err, "Leader Election contract could not be deployed...")

	log.Info(leaderElectionContractAddress.String())

	aliceAddress := sys.cfg.Secrets.Addresses().Alice
	blockNumber := big.NewInt(0)
	isLeader, err := LeaderElectionContract.IsCurrentLeader(&bind.CallOpts{}, aliceAddress, blockNumber)
	require.NoError(t, err)
	require.False(t, isLeader)

	//
	//_, err = LeaderElectionContract.Initialize(opts, aliceAddress, big.NewInt(5))
	//require.NoError(t, err)
	//
	//_, err = LeaderElectionContract.AddParticipant(opts, aliceAddress)
	//require.NoError(t, err)

	//
	////creationBlockNumber, err := LeaderElectionContract.CreationBlockNumber(&bind.CallOpts{})
	////require.NoError(t, err)
	////log.Info("creation block number: %d", creationBlockNumber.String())
	//
	//blockNumber = big.NewInt(1)
	//isLeader, err = LeaderElectionContract.IsCurrentLeader(&bind.CallOpts{}, aliceAddress, blockNumber)
	//require.True(t, isLeader)

}
