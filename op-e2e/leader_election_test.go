package op_e2e

import (
	"context"
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
	NumberOfLeaders := int(cfg.DeployConfig.LeaderElectionNumberOfLeaders)
	sys, accounts, err := startConfigWithTestAccounts(t, &cfg, NumberOfLeaders)
	require.Nil(t, err, "Error starting up system")
	defer sys.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	opts, err := bind.NewKeyedTransactorWithChainID(sys.cfg.Secrets.Alice, cfg.L1ChainIDBig())
	log.Info(opts.GasPrice.String())
	require.Nil(t, err)

	log := testlog.Logger(t, log.LvlInfo)
	log.Info("genesis", "l2", sys.RollupConfig.Genesis.L2, "l1", sys.RollupConfig.Genesis.L1, "l2_time", sys.RollupConfig.Genesis.L2Time)

	l1Client := sys.Clients["l1"]

	// Instantiate the Leader Election Batch Inbox contract
	leaderElectionContractAddress := sys.cfg.L1Deployments.RoundRobinLeaderElection
	log.Info("leaderElectionContractAddress: %s", leaderElectionContractAddress.String())
	leaderElectionContract, err := bindings.NewLeaderElectionBatchInbox(cfg.L1Deployments.RoundRobinLeaderElectionProxy, l1Client)
	require.Nil(t, err)

	// Initialize the Leader Election Batch Inbox contract with the addresses of the Batchers
	batcherAddresses := make([]common.Address, 0, NumberOfLeaders)
	for i := 0; i < NumberOfLeaders; i++ {
		batcherAddresses = append(batcherAddresses, accounts[i].Addr)
	}
	sys.InitLeaderBatchInboxContract(t, batcherAddresses)

	// Check that the leader slots are correctly filled

	NumberOfSlotsPerLeader := int(cfg.DeployConfig.LeaderElectionNumberOfSlotsPerLeader)
	blockNumberOfBatchInboxContractDeployment, err := leaderElectionContract.CreationBlockNumber(&bind.CallOpts{})
	require.Nil(t, err)
	blockNumberOfBatchInboxContractDeploymentInt := int(blockNumberOfBatchInboxContractDeployment.Int64())

	// Check the address of each batcher is assigned to the right leader slot and that it is funded
	expectedBalance := new(big.Int)
	expectedBalance, _ = expectedBalance.SetString("1000000000000000000000000", 10)
	for i := 0; i < NumberOfLeaders; i++ {
		batcherAddress := accounts[i].Addr
		addressBalance, err := l1Client.BalanceAt(ctx, batcherAddress, nil)
		require.NoError(t, err)
		require.Equal(t, expectedBalance, addressBalance, "Batcher address does not seem to be funded.")

		for j := 0; j < NumberOfSlotsPerLeader; j++ {
			blockNumber := blockNumberOfBatchInboxContractDeploymentInt + i*NumberOfSlotsPerLeader + j
			checkIsLeader(t, leaderElectionContract, batcherAddress, big.NewInt(int64(blockNumber)))
		}
	}
}
