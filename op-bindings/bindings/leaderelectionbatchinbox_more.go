// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const LeaderElectionBatchInboxStorageLayoutJSON = "{\"storage\":null,\"types\":{}}"

var LeaderElectionBatchInboxStorageLayout = new(solc.StorageLayout)

var LeaderElectionBatchInboxDeployedBin = "0x"

func init() {
	if err := json.Unmarshal([]byte(LeaderElectionBatchInboxStorageLayoutJSON), LeaderElectionBatchInboxStorageLayout); err != nil {
		panic(err)
	}

	layouts["LeaderElectionBatchInbox"] = LeaderElectionBatchInboxStorageLayout
	deployedBytecodes["LeaderElectionBatchInbox"] = LeaderElectionBatchInboxDeployedBin
}
