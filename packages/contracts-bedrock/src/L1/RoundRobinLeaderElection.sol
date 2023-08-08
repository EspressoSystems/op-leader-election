// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "./LeaderElectionBatchInbox.sol";

/// @title ILeaderElectionBatchInbox
/// @notice Interface for implementing a leader election scheme

contract RoundRobinLeaderElection is ILeaderElectionBatchInbox {

    constructor () public {}

    function isCurrentLeader(address leaderId, uint256 blockNumber) external view returns (bool){
        return false;
    }


    function nextBlocksAsLeader() external view returns (LeaderStatusFlags, uint256, uint256, uint8) {
        return (LeaderStatusFlags.Scheduled, 1,1,1);
    }
}
