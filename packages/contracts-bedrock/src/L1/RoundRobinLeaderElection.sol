// SPDX-License-Identifier: MIT
pragma solidity 0.8.19;

import "./LeaderElectionBatchInbox.sol";

/// @title ILeaderElectionBatchInbox
/// @notice Interface for implementing a leader election scheme

contract RoundRobinLeaderElection is ILeaderElectionBatchInbox {
    address immutable owner;
    uint256 max_number_participants;
    uint32 index_last_inserted_participant;
    mapping(uint256 => address) public participants;
    mapping(address => bool) public is_participant;
    // TODO No need to be public, just for testing purposes. Do this more cleanly
    uint256 public creation_block_number;
    uint8 constant HORIZON = 10; // Number of leader slots that can be checked in advance

    function _uint256FromBytesLittleEndian(uint8[HORIZON] memory _input) private pure returns (uint256) {
        uint256 r = 0;
        for (uint256 i = 0; i < _input.length; i++) {
            r += 2 ** (8 * i) * _input[i];
        }
        return r;
    }

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

    function isCurrentLeader(address _leaderId, uint256 _blockNumber) external view returns (bool) {
        if (_blockNumber < creation_block_number) {
            // if the block number is "too old" then no one is leader
            return false;
        } else {
            uint32 number_participants = index_last_inserted_participant;
            uint256 index_leader = (_blockNumber - creation_block_number) % number_participants;
            return participants[index_leader] == _leaderId;
        }
    }

    function nextBlocksAsLeader() external view returns (LeaderStatusFlags, uint256, uint256, uint8) {
        if (!this.is_participant(msg.sender)) {
            return (LeaderStatusFlags.Invalid, 0, 0, 0);
        } else {
            // Build the bitmap
            uint8[HORIZON] memory leaderSlots;
            for (uint256 i = 0; i < HORIZON; i++) {
                if (this.isCurrentLeader(msg.sender, block.number + i)) {
                    leaderSlots[i] = 1;
                } else {
                    leaderSlots[i] = 0;
                }
            }

            uint256 bitmapAsInt = _uint256FromBytesLittleEndian(leaderSlots);
            return (LeaderStatusFlags.Scheduled, block.number, bitmapAsInt, HORIZON);
        }
    }
}
