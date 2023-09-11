package op_e2e

import (
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum-optimism/optimism/op-bindings/bindings"
	"github.com/ethereum-optimism/optimism/op-node/testlog"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/log"
	"github.com/stretchr/testify/require"
)

func TestLeaderElectionCall(t *testing.T) {
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

func TestLeaderElectionSetup(t *testing.T) {
	InitParallel(t)

	cfg := DefaultSystemConfig(t)

	sys, err := cfg.Start(t)
	require.Nil(t, err, "Error starting up system")
	defer sys.Close()

	opts, err := bind.NewKeyedTransactorWithChainID(sys.cfg.Secrets.Alice, cfg.L1ChainIDBig())
	log.Info(opts.GasPrice.String())
	require.Nil(t, err)

	log := testlog.Logger(t, log.LvlInfo)
	log.Info("genesis", "l2", sys.RollupConfig.Genesis.L2, "l1", sys.RollupConfig.Genesis.L1, "l2_time", sys.RollupConfig.Genesis.L2Time)

	l1Client := sys.Clients["l1"]

	// Initialize the batchers
	// TODO Should this be done in the config? probably yes. Be should be carefully of not breaking things

	// Add the address of the batcher in the leaders' list
	batcherAddress := cfg.Secrets.Addresses().Batcher
	log.Info(batcherAddress.String())

	aliceAddress := cfg.Secrets.Addresses().Alice

	// Instantiante the Leader Election Batch Inbox contract
	leaderElectionContractAddress := sys.cfg.L1Deployments.RoundRobinLeaderElection
	log.Info("leaderElectionContractAddress: %s", leaderElectionContractAddress.String())
	leaderElectionContract, err := bindings.NewLeaderElectionBatchInbox(cfg.L1Deployments.RoundRobinLeaderElectionProxy, l1Client)
	require.Nil(t, err)

	// Add batcher
	tx, err := leaderElectionContract.AddParticipant(opts, batcherAddress)
	require.Nil(t, err)
	require.Nil(t, err, "Adding participant")

	timeout := 10 * time.Duration(cfg.DeployConfig.L1BlockTime) * time.Second
	receipt, err := waitForTransaction(tx.Hash(), l1Client, timeout)
	require.Nil(t, err, "The transaction is sent")
	require.Equal(t, receipt.Status, types.ReceiptStatusSuccessful, "transaction failed")

	// Add Alice
	tx, err = leaderElectionContract.AddParticipant(opts, aliceAddress)
	require.Nil(t, err)
	require.Nil(t, err, "Adding participant")

	receipt, err = waitForTransaction(tx.Hash(), l1Client, timeout)
	require.Nil(t, err, "The transaction is sent")
	require.Equal(t, receipt.Status, types.ReceiptStatusSuccessful, "transaction failed")

	blockNumber := big.NewInt(6)

	// TODO repeat the occurrence of each batcher address N times
	isLeader, err := leaderElectionContract.IsCurrentLeader(&bind.CallOpts{}, batcherAddress, blockNumber)
	require.Nil(t, err)
	require.True(t, isLeader)

}
