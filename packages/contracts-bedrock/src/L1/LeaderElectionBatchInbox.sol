// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/// @title ILeaderElectionBatchInbox
/// @notice Interface for implementing a leader election scheme

abstract contract LeaderElectionBatchInbox {
    enum LeaderStatusFlags {
        Scheduled,
        Unscheduled,
        Invalid
    }

    struct Meta {
        /// Frame metadata
        bytes16 channelId;
        uint16 frameNumber;
        uint32 frameDataLength;
        bool isLast;
        /// The number of L2 blocks in this channel (including the current frame).
        uint16 numL2Blocks;
    }

    /// @notice Allows to submit a batch. This function checks that the caller is the leader for the current block.
    ///
    ///         Checking that the metadata matches the frames must be done in the derivation pipeline. If there is
    ///         a mismatch between the metadata and the frames, the frames must be discarded.
    ///
    ///         Implementations of this contract interface can enforce behaviour based on the metadata submitted
    ///         to this function.
    ///
    /// @param _metas metadata for the frames.
    /// @param _frames frames to be submitted.
    function submit(Meta[] memory _metas, bytes calldata _frames) external {
        // TODO: pass metadata to isCurrentLeader
        bool isLeader = this.isCurrentLeader(msg.sender, block.number);
        require(isLeader, "RoundRobinLeaderElection: submit function must be called by the leader.");
    }

    /// @notice tells if some participant is the leader w.r.t a L1 block number.
    /// @param _leaderId identifier of the leader. If the null address  0x0000000000000000000000000000000000000000 is
    ///         passed, the address of the caller is considered.
    /// @param _blockNumber block number for testing whether some participant is the leader. If blockNumber=0 then the
    ///         block number at the time of the call is considered.
    /// @return true if the leaderId  is the leader w.r.t. block blockNumber, false otherwise.
    function isCurrentLeader(address _leaderId, uint256 _blockNumber) external view virtual returns (bool);

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
