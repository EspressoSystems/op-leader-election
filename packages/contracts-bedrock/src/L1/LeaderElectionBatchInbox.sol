// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/// @title ILeaderElectionBatchInbox
/// @notice Interface for implementing a leader election scheme

abstract contract LeaderElectionBatchInbox {
    /// @notice tells if some participant is the leader w.r.t a L1 block number.
    /// @param _leaderId identifier of the leader. If the null address  0x0000000000000000000000000000000000000000 is
    ///         passed, the address of the caller is considered.
    /// @param _blockNumber block number for testing whether some participant is the leader. If blockNumber=0 then the
    ///         block number at the time of the call is considered.
    /// @return true if the leaderId  is the leader w.r.t. block blockNumber, false otherwise.
    function isCurrentLeader(address _leaderId, uint256 _blockNumber) external view virtual returns (bool);

    enum LeaderStatusFlags {
        Scheduled,
        Unscheduled,
        Invalid
    }

    /// @notice Computes for which blocks from the current one the caller will be the leader in the future.
    /// @param _leaderId identifier of the leader.
    /// @param _blockNumber block number corresponding to the first slot of the returned bitmap array.
    /// @return LeaderStatusFlags: tells whether it is possible to predict when the caller will be the leader in the
    /// future.
    /// @return bool[]: bitmap array encoding leader positions.
    function nextBlocksAsLeader(
        address _leaderId,
        uint256 _blockNumber
    )
        external
        view
        virtual
        returns (LeaderStatusFlags, bool[] memory);
}
