package op_e2e

import (
	"context"
	"math/big"
	"strconv"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/core/types"

	"github.com/ethereum-optimism/optimism/op-node/client"
	"github.com/ethereum-optimism/optimism/op-node/sources"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/ethereum-optimism/optimism/op-service/eth"

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
	log.Info("", "leaderElectionContractAddress", leaderElectionContractAddress.String())
	leaderElectionContract, err := bindings.NewLeaderElectionBatchInbox(sys.cfg.L1Deployments.RoundRobinLeaderElectionProxy, sys.Clients["l1"])
	require.Nil(t, err)
	return leaderElectionContract
}

func getRollupClient(t *testing.T, sys *System) *sources.RollupClient {
	rollupRPCClient, err := rpc.DialContext(context.Background(), sys.RollupNodes["sequencer"].HTTPEndpoint())
	require.Nil(t, err)
	rollupClient := sources.NewRollupClient(client.NewBaseRPCClient(rollupRPCClient))
	return rollupClient
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

// This test covers https://github.com/EspressoSystems/op-leader-election/issues/58 and https://github.com/EspressoSystems/op-leader-election/issues/59
// It instantiates a single batcher (the first one out of three) and creates two L2 blocks that are correctly submitted to L1 by this batcher
func TestLeaderElectionCorrectBatcherSendsTwoBlocks(t *testing.T) {
	InitParallel(t)

	cfg := DefaultSystemConfig(t)

	cfg.switchToV2()

	NumberOfLeaders := int(cfg.DeployConfig.LeaderElectionNumberOfLeaders)
	NumberOfSlotsPerLeader := int(cfg.DeployConfig.LeaderElectionNumberOfSlotsPerLeader)

	log.Info("Deploy configuration:", "Number of leaders", NumberOfLeaders)
	sys, accounts, err := startConfigWithTestAccounts(t, &cfg, NumberOfLeaders)

	require.Nil(t, err, "Error starting up system")
	defer sys.Close()

	sys.InitLeaderBatchInboxContract(t, accounts)

	require.Equal(t, sys.BatchSubmitters[0].Config.BatchInboxVersion, cfg.DeployConfig.InitialBatcherVersion)

	aliceKey := sys.cfg.Secrets.Alice

	l2Client := sys.Clients["sequencer"]

	rollupClient := getRollupClient(t, sys)

	// Start only the first batcher so that we are sure the same batcher is able to publish two L2 blocks
	err = sys.BatchSubmitters[0].Start()
	require.Nil(t, err)

	// Waiting for the batchers to be up
	time.Sleep(5 * time.Second)

	log.Info("Sending transactions to L2...")

	var receipts []*types.Receipt

	numTxs := 2
	for i := 0; i < numTxs; i++ {
		receipt := SendL2Tx(t, cfg, l2Client, aliceKey, func(opts *TxOpts) {
			opts.ToAddr = &cfg.Secrets.Addresses().Bob
			opts.Nonce = uint64(i)
			opts.Value = big.NewInt(1_000)
		})

		receipts = append(receipts, receipt)
		require.NoError(t, err, "Sending L2 tx")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	previousBlockNumber := uint64(0)
	blockNumber := uint64(0)
	for i := 0; i < numTxs; i++ {
		receipt := receipts[i]
		previousBlockNumber = blockNumber
		log.Info("", "previous block number", strconv.Itoa(int(previousBlockNumber)))
		blockNumber = receipt.BlockNumber.Uint64()

		log.Info("", "block number", strconv.Itoa(int(blockNumber)))
		block, _ := l2Client.BlockByNumber(ctx, big.NewInt(int64(blockNumber)))
		log.Info("blockId:  " + eth.ToBlockID(block).String())
		require.NoError(t, waitForSafeHead(ctx, blockNumber, rollupClient))
	}
	// Ensure that the batcher was able to push two consecutive non-empty blocks
	require.True(t, blockNumber-previousBlockNumber > 0)
	require.True(t, (blockNumber-previousBlockNumber) < uint64(NumberOfSlotsPerLeader))
}

func TestLeaderElectionWrongBatcher(t *testing.T) {
	InitParallel(t)

	cfg := DefaultSystemConfig(t)

	cfg.switchToV2()

	NumberOfLeaders := int(cfg.DeployConfig.LeaderElectionNumberOfLeaders)

	log.Info("Deploy configuration:", "Number of leaders", NumberOfLeaders)

	// We generate one extra batcher that will be the wrong one
	numBatchers := NumberOfLeaders + 1
	sys, accounts, err := startConfigWithTestAccounts(t, &cfg, numBatchers)

	require.Nil(t, err, "Error starting up system")
	defer sys.Close()

	require.Equal(t, len(accounts), numBatchers)
	sys.InitLeaderBatchInboxContract(t, accounts)
	require.Equal(t, len(sys.BatchSubmitters), numBatchers)

	aliceKey := sys.cfg.Secrets.Alice

	l2Client := sys.Clients["sequencer"]

	rollupClient := getRollupClient(t, sys)

	// Start the wrong batcher only to prove it cannot produce blocks
	wrongBatcher := sys.BatchSubmitters[numBatchers-1]
	require.Equal(t, wrongBatcher.Config.BatchInboxVersion, cfg.DeployConfig.InitialBatcherVersion)
	err = wrongBatcher.Start()
	require.Nil(t, err)

	// Waiting for the wrong batcher to be up
	time.Sleep(5 * time.Second)

	log.Info("Sending transaction to L2...")

	receipt := SendL2Tx(t, cfg, l2Client, aliceKey, func(opts *TxOpts) {
		opts.ToAddr = &cfg.Secrets.Addresses().Bob
		opts.Value = big.NewInt(1_000)
	})

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	blockNumber := receipt.BlockNumber.Uint64()

	log.Info("", "block number", strconv.Itoa(int(blockNumber)))
	block, _ := l2Client.BlockByNumber(ctx, big.NewInt(int64(blockNumber)))
	log.Info("blockId:  " + eth.ToBlockID(block).String())

	// The block is not produced
	require.Error(t, waitForSafeHead(ctx, blockNumber, rollupClient))

}

func TestCorrectSequenceOfBatchersFourEpochs(t *testing.T) {
	InitParallel(t)

	cfg := DefaultSystemConfig(t)

	cfg.switchToV2()

	NumberOfLeaders := int(cfg.DeployConfig.LeaderElectionNumberOfLeaders)
	NumberOfSlotsPerLeader := int(cfg.DeployConfig.LeaderElectionNumberOfSlotsPerLeader)
	log.Info("Deploy configuration:", "Number of leaders", NumberOfLeaders)
	sys, accounts, err := startConfigWithTestAccounts(t, &cfg, NumberOfLeaders)

	require.Nil(t, err, "Error starting up system")
	defer sys.Close()

	sys.InitLeaderBatchInboxContract(t, accounts)

	require.Equal(t, sys.BatchSubmitters[0].Config.BatchInboxVersion, cfg.DeployConfig.InitialBatcherVersion)

	aliceKey := sys.cfg.Secrets.Alice

	l2Client := sys.Clients["sequencer"]

	rollupClient := getRollupClient(t, sys)

	// Start all the batchers
	for i := 0; i < NumberOfLeaders; i++ {
		err = sys.BatchSubmitters[i].Start()
		require.Nil(t, err)
	}
	// Waiting for the batchers to be up
	time.Sleep(5 * time.Second)

	log.Info("Sending transactions to L2...")

	var receipts []*types.Receipt

	numTxs := 7 * NumberOfSlotsPerLeader
	for i := 0; i < numTxs; i++ {
		receipt := SendL2Tx(t, cfg, l2Client, aliceKey, func(opts *TxOpts) {
			opts.ToAddr = &cfg.Secrets.Addresses().Bob
			opts.Nonce = uint64(i)
			opts.Value = big.NewInt(1_000)
		})

		receipts = append(receipts, receipt)
		require.NoError(t, err, "Sending L2 tx")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	minBlockNumber := uint64(1000) // Very high value to start with
	maxBlockNumber := uint64(0)
	for i := 0; i < numTxs; i++ {
		receipt := receipts[i]

		blockNumber := receipt.BlockNumber.Uint64()
		// Update min and max block numbers
		if blockNumber > maxBlockNumber {
			maxBlockNumber = blockNumber
		}

		if blockNumber < minBlockNumber {
			minBlockNumber = blockNumber
		}

		log.Info("", "block number", strconv.Itoa(int(blockNumber)))
		block, _ := l2Client.BlockByNumber(ctx, big.NewInt(int64(blockNumber)))
		log.Info("blockId:  " + eth.ToBlockID(block).String())
		require.NoError(t, waitForSafeHead(ctx, blockNumber, rollupClient))
	}
	// The gap between the min and max block number should cover 4 epochs at least
	log.Info("", "maxBlockNumber-minBlockNumber", maxBlockNumber-minBlockNumber)
	require.True(t, maxBlockNumber-minBlockNumber >= 4*uint64(NumberOfSlotsPerLeader))

}

// We produce several blocks, everything should go through despite the presence of a wrong batcher and the absence of a good one.
func TestMixOfGoodAndBadBatchers(t *testing.T) {
	InitParallel(t)

	cfg := DefaultSystemConfig(t)

	cfg.switchToV2()

	NumberOfLeaders := int(cfg.DeployConfig.LeaderElectionNumberOfLeaders)
	NumberOfSlotsPerLeader := int(cfg.DeployConfig.LeaderElectionNumberOfSlotsPerLeader)

	log.Info("Deploy configuration:", "Number of leaders", NumberOfLeaders)
	// We generate one extra batcher that will be the wrong one
	numBatchers := NumberOfLeaders + 1

	sys, accounts, err := startConfigWithTestAccounts(t, &cfg, numBatchers)

	require.Nil(t, err, "Error starting up system")
	defer sys.Close()

	sys.InitLeaderBatchInboxContract(t, accounts)

	require.Equal(t, sys.BatchSubmitters[0].Config.BatchInboxVersion, cfg.DeployConfig.InitialBatcherVersion)

	aliceKey := sys.cfg.Secrets.Alice

	l2Client := sys.Clients["sequencer"]

	rollupClient := getRollupClient(t, sys)

	// Start all the batchers but replace the second one (which is good) by the bad one
	for i := 0; i < NumberOfLeaders; i++ {
		if i == 1 { // Start the wrong batcher
			err = sys.BatchSubmitters[numBatchers-1].Start()
		} else {
			err = sys.BatchSubmitters[i].Start()
		}
		require.Nil(t, err)
	}
	// Waiting for the batchers to be up
	time.Sleep(5 * time.Second)

	log.Info("Sending transactions to L2...")

	var receipts []*types.Receipt

	numTxs := 7 * NumberOfSlotsPerLeader
	for i := 0; i < numTxs; i++ {
		receipt := SendL2Tx(t, cfg, l2Client, aliceKey, func(opts *TxOpts) {
			opts.ToAddr = &cfg.Secrets.Addresses().Bob
			opts.Nonce = uint64(i)
			opts.Value = big.NewInt(1_000)
		})

		receipts = append(receipts, receipt)
		require.NoError(t, err, "Sending L2 tx")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// All the blocks should be processed correctly
	for i := 0; i < numTxs; i++ {
		receipt := receipts[i]

		blockNumber := receipt.BlockNumber.Uint64()

		log.Info("", "block number", strconv.Itoa(int(blockNumber)))
		block, _ := l2Client.BlockByNumber(ctx, big.NewInt(int64(blockNumber)))
		log.Info("blockId:  " + eth.ToBlockID(block).String())
		require.NoError(t, waitForSafeHead(ctx, blockNumber, rollupClient))
	}
}

// We produce several blocks, everything should go through despite the absence of a good batcher
func TestMissingGoodBatcher(t *testing.T) {
	InitParallel(t)

	cfg := DefaultSystemConfig(t)

	cfg.switchToV2()

	NumberOfLeaders := int(cfg.DeployConfig.LeaderElectionNumberOfLeaders)
	NumberOfSlotsPerLeader := int(cfg.DeployConfig.LeaderElectionNumberOfSlotsPerLeader)
	log.Info("Deploy configuration:", "Number of leaders", NumberOfLeaders)

	sys, accounts, err := startConfigWithTestAccounts(t, &cfg, NumberOfLeaders)

	require.Nil(t, err, "Error starting up system")
	defer sys.Close()

	sys.InitLeaderBatchInboxContract(t, accounts)

	require.Equal(t, sys.BatchSubmitters[0].Config.BatchInboxVersion, cfg.DeployConfig.InitialBatcherVersion)

	aliceKey := sys.cfg.Secrets.Alice

	l2Client := sys.Clients["sequencer"]

	rollupClient := getRollupClient(t, sys)

	// Start all the batchers but one
	for i := 0; i < NumberOfLeaders; i++ {
		if i != 1 { // Skip one good batcher
			err = sys.BatchSubmitters[i].Start()
		}
		require.Nil(t, err)
	}
	// Waiting for the batchers to be up
	time.Sleep(5 * time.Second)

	log.Info("Sending transactions to L2...")

	var receipts []*types.Receipt

	numTxs := 7 * NumberOfSlotsPerLeader
	for i := 0; i < numTxs; i++ {
		receipt := SendL2Tx(t, cfg, l2Client, aliceKey, func(opts *TxOpts) {
			opts.ToAddr = &cfg.Secrets.Addresses().Bob
			opts.Nonce = uint64(i)
			opts.Value = big.NewInt(1_000)
		})

		receipts = append(receipts, receipt)
		require.NoError(t, err, "Sending L2 tx")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// All the blocks should be processed correctly
	for i := 0; i < numTxs; i++ {
		receipt := receipts[i]

		blockNumber := receipt.BlockNumber.Uint64()
		log.Info("", "block number", strconv.Itoa(int(blockNumber)))
		block, _ := l2Client.BlockByNumber(ctx, big.NewInt(int64(blockNumber)))
		log.Info("blockId:  " + eth.ToBlockID(block).String())
		require.NoError(t, waitForSafeHead(ctx, blockNumber, rollupClient))
	}
}
