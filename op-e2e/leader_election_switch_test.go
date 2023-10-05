package op_e2e

import (
	"context"
	"math/big"
	"strconv"
	"testing"
	"time"

	"github.com/ethereum-optimism/optimism/op-service/eth"

	"github.com/ethereum/go-ethereum/log"
	"github.com/stretchr/testify/require"
)

func TestLeaderElectionSwitchBatcherFromV1ToV2(t *testing.T) {
	// InitParallel(t)

	cfg := defaultConfigLeaderElection(t)

	NumberOfLeaders := int(cfg.DeployConfig.LeaderElectionNumberOfLeaders)
	log.Info("Deploy configuration:", "Number of leaders", NumberOfLeaders)
	sys, accounts, err := startConfigWithTestAccounts(t, &cfg, NumberOfLeaders)

	require.Nil(t, err, "Error starting up system")
	defer sys.Close()

	sys.InitLeaderBatchInboxContract(t, accounts)

	aliceKey := sys.cfg.Secrets.Alice

	l2Client := sys.Clients["sequencer"]

	rollupClient := getRollupClient(t, sys)

	// Start all batchers
	for i := 0; i < NumberOfLeaders; i++ {
		err = sys.BatchSubmitters[i].Start()
		require.Nil(t, err)
	}

	// Waiting for the batchers to be up
	time.Sleep(5 * time.Second)

	{
		log.Info("Sending a transaction to L2...")

		receipt := SendL2Tx(t, cfg, l2Client, aliceKey, func(opts *TxOpts) {
			opts.ToAddr = &cfg.Secrets.Addresses().Bob
			opts.Value = big.NewInt(1_000)
			opts.Nonce = uint64(0)
		})
		require.NoError(t, err, "Sending L2 tx")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		blockNumber := receipt.BlockNumber.Uint64()
		log.Info("", "block receipt", strconv.Itoa(int(blockNumber)))
		block, _ := l2Client.BlockByNumber(ctx, big.NewInt(int64(blockNumber)))
		log.Info("blockId:  " + eth.ToBlockID(block).String())

		require.NoError(t, waitForSafeHead(ctx, receipt.BlockNumber.Uint64(), rollupClient))

	}

	log.Info("Setting SystemConfig BatcherHash to V2")

	sys.SetBatchInboxToV2(t)

	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		log.Info("Sending another transaction to L2...", "counter", i)

		receipt := SendL2Tx(t, cfg, l2Client, aliceKey, func(opts *TxOpts) {
			opts.ToAddr = &cfg.Secrets.Addresses().Bob
			opts.Value = big.NewInt(1_000)
			opts.Nonce = uint64(i + 1)
		})
		require.NoError(t, err, "Sending L2 tx")

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		blockNumber := receipt.BlockNumber.Uint64()
		log.Info("", "block receipt", strconv.Itoa(int(blockNumber)))
		block, _ := l2Client.BlockByNumber(ctx, big.NewInt(int64(blockNumber)))
		log.Info("blockId:  " + eth.ToBlockID(block).String())

		require.NoError(t, waitForSafeHead(ctx, receipt.BlockNumber.Uint64(), rollupClient))

	}
}
