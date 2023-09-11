package op_e2e

import (
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

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

func addNewLeader(
	t *testing.T,
	timeout time.Duration,
	l1Client *ethclient.Client,
	address common.Address,
	contract *bindings.LeaderElectionBatchInbox,
	opts *bind.TransactOpts) {

	tx, err := contract.AddParticipant(opts, address)
	require.Nil(t, err)
	require.Nil(t, err, "Adding participant")

	receipt, err := waitForTransaction(tx.Hash(), l1Client, timeout)
	require.Nil(t, err, "The transaction is sent")
	require.Equal(t, receipt.Status, types.ReceiptStatusSuccessful, "transaction failed")
}

func checkIsLeader(
	t *testing.T,
	contract *bindings.LeaderElectionBatchInbox,
	address common.Address,
	blockNumber *big.Int) {
	isLeader, err := contract.IsCurrentLeader(&bind.CallOpts{}, address, blockNumber)
	require.Nil(t, err)
	require.True(t, isLeader)
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

	// Instantiate the Leader Election Batch Inbox contract
	leaderElectionContractAddress := sys.cfg.L1Deployments.RoundRobinLeaderElection
	log.Info("leaderElectionContractAddress: %s", leaderElectionContractAddress.String())
	leaderElectionContract, err := bindings.NewLeaderElectionBatchInbox(cfg.L1Deployments.RoundRobinLeaderElectionProxy, l1Client)
	require.Nil(t, err)

	timeout := 10 * time.Duration(cfg.DeployConfig.L1BlockTime) * time.Second

	// Add batcher
	addNewLeader(t, timeout, l1Client, batcherAddress, leaderElectionContract, opts)

	// Add Alice
	addNewLeader(t, timeout, l1Client, aliceAddress, leaderElectionContract, opts)

	// TODO repeat leader occurrences N times

	// Check leader slots are correctly filled
	checkIsLeader(t, leaderElectionContract, batcherAddress, big.NewInt(4))
	checkIsLeader(t, leaderElectionContract, aliceAddress, big.NewInt(5))
	checkIsLeader(t, leaderElectionContract, batcherAddress, big.NewInt(6))
	checkIsLeader(t, leaderElectionContract, aliceAddress, big.NewInt(7))

}
