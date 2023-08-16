// SPDX-License-Identifier: MIT
pragma solidity 0.8.19;

import "./LeaderElectionBatchInbox.sol";

/// @title ILeaderElectionBatchInbox
/// @notice Interface for implementing a leader election scheme

contract RoundRobinLeaderElection is LeaderElectionBatchInbox {
    address immutable owner;
    uint256 max_number_participants;
    uint32 index_last_inserted_participant;
    mapping(uint256 => address) public participants;
    mapping(address => bool) public is_participant;
    // TODO No need to be public, just for testing purposes. Do this more cleanly
    uint256 public creation_block_number;
    uint8 constant HORIZON = 10; // Number of leader slots that can be checked in advance

    constructor(uint256 _n) {
        owner = msg.sender;
        max_number_participants = _n;
        creation_block_number = block.number;
    }

    function addParticipant(address _addr) public {
        require(
            msg.sender == owner, "RoundRobinLeaderElection: only the creator of this contract can call this function."
        );
        require(
            index_last_inserted_participant < max_number_participants,
            "RoundRobinLeaderElection: list of participants is full."
        );

        participants[index_last_inserted_participant] = _addr;
        is_participant[_addr] = true;
        index_last_inserted_participant++;
    }

    function isCurrentLeader(address _leaderId, uint256 _blockNumber) external view override returns (bool) {
        if (_blockNumber < creation_block_number) {
            // if the block number is "too old" then no one is leader
            return false;
        } else {
            uint32 number_participants = index_last_inserted_participant;
            uint256 index_leader = (_blockNumber - creation_block_number) % number_participants;
            return participants[index_leader] == _leaderId;
        }
    }

    function nextBlocksAsLeader(
        address _leaderId,
        uint256 _blockNumber
    )
        external
        view
        override
        returns (LeaderStatusFlags, bool[] memory)
    {
        if (!this.is_participant(_leaderId)) {
            bool[] memory emptyArray;
            return (LeaderStatusFlags.Invalid, emptyArray);
        } else {
            // Build the bitmap
            bool[] memory leaderSlots = new bool[](HORIZON);
            for (uint256 i = 0; i < HORIZON; i++) {
                leaderSlots[i] = this.isCurrentLeader(_leaderId, _blockNumber + i);
            }

            return (LeaderStatusFlags.Scheduled, leaderSlots);
        }
    }
}
