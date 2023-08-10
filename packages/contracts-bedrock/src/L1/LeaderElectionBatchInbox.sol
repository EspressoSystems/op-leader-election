// SPDX-License-Identifier: MIT
pragma solidity 0.8.19;

/// @title ILeaderElectionBatchInbox
/// @notice Interface for implementing a leader election scheme

interface ILeaderElectionBatchInbox {
    /// @notice tells if some participant is the leader w.r.t a L1 block number.
    /// @param leaderId identifier of the leader. If the null address  0x0000000000000000000000000000000000000000 is passed, the address of the caller is considered.
    /// @param blockNumber block number for testing whether some participant is the leader. If blockNumber=0 then the block number at the time of the call is considered.
    /// @return true if the leaderId  is the leader w.r.t. block blockNumber, false otherwise.
    function isCurrentLeader(address leaderId, uint256 blockNumber) external view returns (bool);

    enum LeaderStatusFlags {
        Scheduled,
        Unscheduled,
        Invalid
    }

    /// @notice Computes for which blocks from the current one the caller will be the leader in the future.
    ///         Example if the current block is 50, and the caller is the leader for blocks 50, 51, 53, 57, then the integer will represent the bitmap array [1,1,0,1,0,0,0,1], encoded as 139 in little endian.
    ///         Additionally we provide another uint8 value to specify that the 0s after some position mean "not defined yet" instead of "not a leader". In this example this value is N=8.
    /// @return LeaderStatusFlags: tells whether it is possible to predict when the caller will be the leader in the future.
    /// @return uint256: current block number.
    /// @return uin256: represents a bitmap array encoding leader positions.
    /// @return uint8: integer N that specifies that only values in the slice [0,N-1] from the bitmap array is relevant.
    function nextBlocksAsLeader() external view returns (LeaderStatusFlags, uint256, uint256, uint8);
}
