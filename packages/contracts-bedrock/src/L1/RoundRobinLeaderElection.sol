// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import { OwnableUpgradeable } from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import { Semver } from "../universal/Semver.sol";
import { LeaderElectionBatchInbox } from "./LeaderElectionBatchInbox.sol";

/// @title ILeaderElectionBatchInbox
/// @notice Interface for implementing a leader election scheme

contract RoundRobinLeaderElection is LeaderElectionBatchInbox, OwnableUpgradeable, Semver {
    // @notice The number of leader slots that can be checked in advance
    uint8 public constant HORIZON = 10;

    uint256 public max_number_participants;
    uint32 index_last_inserted_participant;
    mapping(uint256 => address) public participants;
    mapping(address => bool) public is_participant;
    // TODO No need to be public, just for testing purposes. Do this more cleanly

    constructor() Semver(0, 1, 0) {
        initialize({ _owner: address(0xdEaD), _max_number_participants: 0 });
    }

    function initialize(address _owner, uint256 _max_number_participants) public reinitializer(2) {
        __Ownable_init();
        transferOwnership(_owner);

        max_number_participants = _max_number_participants;
        creation_block_number = block.number;
    }

    // TODO for production purposes it might be desirable to make this function "onlyOwner". However the tests would be
    // harder to write. See ticket https://github.com/EspressoSystems/op-leader-election/issues/73
    function addParticipant(address _addr) public override {
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
