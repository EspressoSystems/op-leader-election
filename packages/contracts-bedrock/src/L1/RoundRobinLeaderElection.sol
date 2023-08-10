// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "./LeaderElectionBatchInbox.sol";

/// @title ILeaderElectionBatchInbox
/// @notice Interface for implementing a leader election scheme

contract RoundRobinLeaderElection is ILeaderElectionBatchInbox {
    address immutable owner;
    uint256 max_number_participants;
    uint256 index_last_inserted_participant;
    mapping(uint256 => address) public participants;
    mapping(address => bool) public is_participant;
    // TODO No need to be public, just for testing purposes. Do this more cleanly
    uint256 public creation_block_number;
    uint8 constant HORIZON = 10; // Number of leader slots that can be checked in advance

    // TODO gas optimization
    function _uint256FromBytesLittleEndian(uint8[HORIZON] memory input) private pure returns (uint256) {
        uint256 r = 0;
        for (uint256 i = 0; i < input.length; i++) {
            r += 2 ** (8 * i) * input[i];
        }
        return r;
    }

    constructor(uint256 n) {
        owner = msg.sender;
        max_number_participants = n;
        creation_block_number = block.number;
    }

    function addParticipant(address addr) public {
        // TODO write this more cleanly
        if (msg.sender != owner) {
            revert("Only the creator of this contract can call this function.");
        }
        if (index_last_inserted_participant == max_number_participants) {
            revert("List of participants is full.");
        }

        participants[index_last_inserted_participant] = addr;
        is_participant[addr] = true;
        index_last_inserted_participant++;
    }

    function isCurrentLeader(address leaderId, uint256 blockNumber) external view returns (bool) {
        if (blockNumber < creation_block_number) {
            // if the block number is "too old" then no one is leader
            return false;
        } else {
            uint256 index_leader = (blockNumber - creation_block_number) % max_number_participants;
            return participants[index_leader] == leaderId;
        }
    }

    function nextBlocksAsLeader() external view returns (LeaderStatusFlags, uint256, uint256, uint8) {
        if (!this.is_participant(msg.sender)) {
            return (LeaderStatusFlags.Invalid, 0, 0, 0);
        } else {
            // Build the bitmap
            uint8[HORIZON] memory leaderSlots;
            for (uint256 i = 0; i < HORIZON; i++) {
                if (this.isCurrentLeader(msg.sender, block.number)) {
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
