// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const LeaderElectionBatchInboxStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/L1/LeaderElectionBatchInbox.sol:LeaderElectionBatchInbox\",\"label\":\"creation_block_number\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint256\"}],\"types\":{\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"}}}"

var LeaderElectionBatchInboxStorageLayout = new(solc.StorageLayout)

var LeaderElectionBatchInboxDeployedBin = "0x"

func init() {
	if err := json.Unmarshal([]byte(LeaderElectionBatchInboxStorageLayoutJSON), LeaderElectionBatchInboxStorageLayout); err != nil {
		panic(err)
	}

	layouts["LeaderElectionBatchInbox"] = LeaderElectionBatchInboxStorageLayout
	deployedBytecodes["LeaderElectionBatchInbox"] = LeaderElectionBatchInboxDeployedBin
}
