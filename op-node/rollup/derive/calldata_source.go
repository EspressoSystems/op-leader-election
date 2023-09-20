package derive

import (
	"context"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"

	"github.com/ethereum-optimism/optimism/op-node/rollup"
	"github.com/ethereum-optimism/optimism/op-service/eth"
)

type DataIter interface {
	Next(ctx context.Context) (eth.Data, error)
}

type L1TransactionFetcher interface {
	InfoAndTxsByHash(ctx context.Context, hash common.Hash) (eth.BlockInfo, types.Transactions, error)
	L1ReceiptsFetcher
}

// DataSourceFactory readers raw transactions from a given block & then filters for
// batch submitter transactions.
// This is not a stage in the pipeline, but a wrapper for another stage in the pipeline
type DataSourceFactory struct {
	log     log.Logger
	cfg     *rollup.Config
	fetcher L1TransactionFetcher
}

func NewDataSourceFactory(log log.Logger, cfg *rollup.Config, fetcher L1TransactionFetcher) *DataSourceFactory {
	return &DataSourceFactory{log: log, cfg: cfg, fetcher: fetcher}
}

// OpenData returns a DataIter. This struct implements the `Next` function.
func (ds *DataSourceFactory) OpenData(ctx context.Context, id eth.BlockID, batcherAddr common.Address, batcherHashVersion uint8) DataIter {
	return NewDataSource(ctx, ds.log, ds.cfg, ds.fetcher, id, batcherAddr, batcherHashVersion)
}

// DataSource is a fault tolerant approach to fetching data.
// The constructor will never fail & it will instead re-attempt the fetcher
// at a later point.
type DataSource struct {
	// Internal state + data
	open bool
	data []eth.Data
	// Required to re-attempt fetching
	id      eth.BlockID
	cfg     *rollup.Config // TODO: `DataFromEVMTransactions` should probably not take the full config
	fetcher L1TransactionFetcher
	log     log.Logger

	batcherAddr        common.Address
	batcherHashVersion uint8
}

// NewDataSource creates a new calldata source. It suppresses errors in fetching the L1 block if they occur.
// If there is an error, it will attempt to fetch the result on the next call to `Next`.
func NewDataSource(ctx context.Context, log log.Logger, cfg *rollup.Config, fetcher L1TransactionFetcher, block eth.BlockID, batcherAddr common.Address, batcherHashVersion uint8) DataIter {
	_, txs, err := fetcher.InfoAndTxsByHash(ctx, block.Hash)
	if err != nil {
		return &DataSource{
			open:               false,
			id:                 block,
			cfg:                cfg,
			fetcher:            fetcher,
			log:                log,
			batcherAddr:        batcherAddr,
			batcherHashVersion: batcherHashVersion,
		}
	}
	_, receipts, err := fetcher.FetchReceipts(ctx, block.Hash)
	if err != nil {
		return &DataSource{
			open:               false,
			id:                 block,
			cfg:                cfg,
			fetcher:            fetcher,
			log:                log,
			batcherAddr:        batcherAddr,
			batcherHashVersion: batcherHashVersion,
		}
	} else {
		if batcherHashVersion == 0 {
			return &DataSource{
				open: true,
				data: DataFromEVMTransactions(cfg, batcherAddr, txs, log.New("origin", block)),
			}
		} else {
			return &DataSource{
				open: true,
				data: DataFromEVMTransactionsV2(cfg, txs, receipts, log.New("origin", block)),
			}
		}
	}
}

// Next returns the next piece of data if it has it. If the constructor failed, this
// will attempt to reinitialize itself. If it cannot find the block it returns a ResetError
// otherwise it returns a temporary error if fetching the block returns an error.
func (ds *DataSource) Next(ctx context.Context) (eth.Data, error) {
	if !ds.open {
		if _, txs, err := ds.fetcher.InfoAndTxsByHash(ctx, ds.id.Hash); err == nil {
			if ds.batcherHashVersion == 0 {
				ds.open = true
				ds.data = DataFromEVMTransactions(ds.cfg, ds.batcherAddr, txs, log.New("origin", ds.id))
			} else {
				// TODO FetchReceipts also calls InfoAndTxsByHash, so the txs could
				// also be returned from that call. This would save a round trip.
				if _, receipts, err := ds.fetcher.FetchReceipts(ctx, ds.id.Hash); err == nil {
					ds.open = true
					ds.data = DataFromEVMTransactionsV2(ds.cfg, txs, receipts, log.New("origin", ds.id))
					// TODO: handle errors in a single place, if possible.
				} else if errors.Is(err, ethereum.NotFound) {
					return nil, NewResetError(fmt.Errorf("failed to open calldata source: %w", err))
				} else {
					return nil, NewTemporaryError(fmt.Errorf("failed to open calldata source: %w", err))
				}
			}
		} else if errors.Is(err, ethereum.NotFound) {
			return nil, NewResetError(fmt.Errorf("failed to open calldata source: %w", err))
		} else {
			return nil, NewTemporaryError(fmt.Errorf("failed to open calldata source: %w", err))
		}
	}
	if len(ds.data) == 0 {
		return nil, io.EOF
	} else {
		data := ds.data[0]
		ds.data = ds.data[1:]
		return data, nil
	}
}

// DataFromEVMTransactions filters all of the transactions and returns the calldata from transactions
// that are sent to the batch inbox address from the batch sender address.
// This will return an empty array if no valid transactions are found.
func DataFromEVMTransactions(config *rollup.Config, batcherAddr common.Address, txs types.Transactions, log log.Logger) []eth.Data {
	var out []eth.Data
	l1Signer := config.L1Signer()
	for j, tx := range txs {
		if to := tx.To(); to != nil && *to == config.BatchInboxAddress {
			seqDataSubmitter, err := l1Signer.Sender(tx) // optimization: only derive sender if To is correct
			if err != nil {
				log.Warn("tx in inbox with invalid signature", "index", j, "err", err)
				continue // bad signature, ignore
			}
			// some random L1 user might have sent a transaction to our batch inbox, ignore them
			if seqDataSubmitter != batcherAddr {
				log.Warn("tx in inbox with unauthorized submitter", "index", j, "err", err)
				continue // not an authorized batch submitter, ignore
			}
			out = append(out, tx.Data())
		}
	}
	return out
}

// DataFromEVMTransactionsV2 filters all of the transactions and returns the
// calldata from transactions that are sent to the submit function of the batch
// inbox contract and did not revert.
// This will return an empty array if no valid transactions are found.
func DataFromEVMTransactionsV2(config *rollup.Config, txs types.Transactions, receipts types.Receipts, log log.Logger) []eth.Data {

	msg := "Entering DataFromEVMTransactionsV2... with " + strconv.Itoa(txs.Len()) + " transactions"
	log.Info(msg)
	msg = "config.BatchInboxContractAddr: " + config.BatchInboxAddress.String()
	log.Info(msg)
	var out []eth.Data
	for j, tx := range txs {
		msg := "tx.To()= " + tx.To().String()
		log.Info(msg)
		msg = "config.BatchInboxContractAddr: " + config.BatchInboxContractAddr.String()
		log.Info(msg)
		if to := tx.To(); to != nil && *to == config.BatchInboxContractAddr {

			log.Info("Trying to process L1 transactions...")

			receipt := receipts[j]
			data := tx.Data()
			transactionHeader := data[:4]
			log.Warn("len(data) " + strconv.Itoa(len(data)))
			log.Info("transactionHeader: " + string(transactionHeader))
			log.Info("transaction data: " + string(data))
			log.Info("SubmitAbi.ID: " + string(SubmitAbi.ID))

			// Exclude transactions if L1 transaction did not call submit function.
			// TODO uncomment
			//if len(data) < 4 || !bytes.Equal(data[:4], SubmitAbi.ID) {
			//	log.Warn("tx sent to inbox contract did not call submit function", "index", j)
			//	continue // not calling submit function, ignore
			//}

			// Exclude transactions if L1 transaction reverted.
			if receipt.Status != types.ReceiptStatusSuccessful {
				log.Warn("tx sent to inbox contract reverted", "index", j)
				continue // reverted, ignore
			}
			log.Info("Data transaction appended:" + string(data))
			out = append(out, data)
		}
	}
	return out
}
