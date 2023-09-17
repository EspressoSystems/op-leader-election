package op_e2e

import (
	"context"
	"github.com/ethereum-optimism/optimism/op-node/client"
	"github.com/ethereum-optimism/optimism/op-node/rollup/derive"
	"github.com/ethereum-optimism/optimism/op-node/sources"
	"github.com/ethereum/go-ethereum/rpc"
	"math/big"
	"strconv"
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

func getBatchInboxContract(t *testing.T, sys *System) *bindings.LeaderElectionBatchInbox {
	// Instantiate the Leader Election Batch Inbox contract
	leaderElectionContractAddress := sys.cfg.L1Deployments.RoundRobinLeaderElection
	log.Info("leaderElectionContractAddress: %s", leaderElectionContractAddress.String())
	leaderElectionContract, err := bindings.NewLeaderElectionBatchInbox(sys.cfg.L1Deployments.RoundRobinLeaderElectionProxy, sys.Clients["l1"])
	require.Nil(t, err)
	return leaderElectionContract
}

func TestLeaderElectionSetup(t *testing.T) {
	InitParallel(t)

	// TODO extract function to generate the setup. Problem, by doing so one gets some weird error related to the Batch inbox contract bindings
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
	leaderElectionContract := getBatchInboxContract(t, sys)

	sys.InitLeaderBatchInboxContract(t, accounts)

	NumberOfSlotsPerLeader := int(cfg.DeployConfig.LeaderElectionNumberOfSlotsPerLeader)
	blockNumberOfBatchInboxContractDeployment, err := leaderElectionContract.CreationBlockNumber(&bind.CallOpts{})
	require.Nil(t, err)
	blockNumberOfBatchInboxContractDeploymentInt := int(blockNumberOfBatchInboxContractDeployment.Int64())

	// Check the address of each batcher is assigned to the right leader slot and that it is funded
	expectedBalance := new(big.Int)
	expectedBalance, _ = expectedBalance.SetString("1000000000000000000000000", 10)
	for i := 0; i < NumberOfLeaders; i++ {
		batcherAddress := sys.BatchSubmitters[i].TxManager.From()
		addressBalance, err := l1Client.BalanceAt(ctx, batcherAddress, nil)
		require.NoError(t, err)
		require.Equal(t, expectedBalance, addressBalance, "Batcher address does not seem to be funded.")

		for j := 0; j < NumberOfSlotsPerLeader; j++ {
			blockNumber := blockNumberOfBatchInboxContractDeploymentInt + i*NumberOfSlotsPerLeader + j
			checkIsLeader(t, leaderElectionContract, batcherAddress, big.NewInt(int64(blockNumber)))
		}
	}
}

func TestLeaderElectionCorrectBatcherSendOneBlock(t *testing.T) {
	InitParallel(t)

	// TODO extract function to generate the setup(see function TestLe
	cfg := DefaultSystemConfig(t)
	cfg.DeployConfig.InitialBatcherVersion = derive.BatchV2Type // TODO Make a function for  that
	NumberOfLeaders := int(cfg.DeployConfig.LeaderElectionNumberOfLeaders)
	sys, accounts, err := startConfigWithTestAccounts(t, &cfg, NumberOfLeaders)

	require.Nil(t, err, "Error starting up system")
	defer sys.Close()

	sys.InitLeaderBatchInboxContract(t, accounts)
	ctx := context.Background()
	aliceKey := sys.cfg.Secrets.Alice

	l1Client := sys.Clients["l1"]
	l2CLient := sys.Clients["sequencer"]

	rollupRPCClient, err := rpc.DialContext(context.Background(), sys.RollupNodes["sequencer"].HTTPEndpoint())
	require.Nil(t, err)
	rollupClient := sources.NewRollupClient(client.NewBaseRPCClient(rollupRPCClient))

	for i := 0; i < NumberOfLeaders; i++ {
		err = sys.BatchSubmitters[i].Start()
		require.Nil(t, err)
	}

	// Waiting for the batchers to be up
	time.Sleep(5 * time.Second)

	log.Info("Sending a transaction to L2...")
	receipt := SendL2Tx(t, cfg, l2CLient, aliceKey, func(opts *TxOpts) {
		opts.ToAddr = &cfg.Secrets.Addresses().Bob
		opts.Value = big.NewInt(1_000)
	})
	require.NoError(t, waitForSafeHead(ctx, receipt.BlockNumber.Uint64(), rollupClient))

	latestBlock := latestBlock(t, l1Client)
	log.Info("Latest block ", strconv.Itoa(int(latestBlock)))

}
