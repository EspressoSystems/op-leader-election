// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

/// @title LeaderElectionBatchInbox
/// @notice ...

interface ILeaderElectionBatchInbox {

    /// @notice tells if the caller of the function is the leader for the current block.
    /// @return true if the caller is the leader, false otherwise.
    function isCurrentLeader() external view returns (bool);

    enum LeaderStatusFlags { Scheduled, Unscheduled, Invalid }

    /// @notice Computes for which blocks from the current one the caller will be the leader in the future.
    ///         Example if the current block is 50, and the caller is the leader for blocks 50, 51, 53, 57, then the first integer will represent the bitmap array [1,1,0,1,0,0,0,1], encoded as 139 in little endian.
    /// @return LeaderStatusFlags: tells whether it is possible to predict when the caller will be the leader in the future.
    /// @return int: bitmap array encoding when the caller will lead after the current block.
    function nextBlocksAsLeader() external view returns (LeaderStatusFlags, int);
}
