// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import { OwnableUpgradeable } from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import { Semver } from "../universal/Semver.sol";
import { LeaderElectionBatchInbox } from "./LeaderElectionBatchInbox.sol";

/// @title ILeaderElectionBatchInbox
/// @notice Interface for implementing a leader election scheme

contract RoundRobinLeaderElection is LeaderElectionBatchInbox, OwnableUpgradeable, Semver {
    // @notice The number of leader slots that can be checked in advance
    uint256 public maxNumberParticipants;
    uint256 public numberOfSlotsPerLeader;
    uint256 public horizon;
    uint32 private indexLastInsertedParticipant;

    mapping(uint256 => address) public participants;
    mapping(address => bool) public isParticipant;
    // TODO No need to be public, just for testing purposes. Do this more cleanly

    constructor() Semver(0, 1, 0) {
        initialize({ _owner: address(0xdEaD), _maxNumberParticipants: 0, _numberOfSlotsPerLeader: 0 });
    }

    function initialize(
        address _owner,
        uint256 _maxNumberParticipants,
        uint256 _numberOfSlotsPerLeader
    )
        public
        reinitializer(2)
    {
        __Ownable_init();
        transferOwnership(_owner);

        maxNumberParticipants = _maxNumberParticipants;
        numberOfSlotsPerLeader = _numberOfSlotsPerLeader;
        creationBlockNumber = block.number;
        horizon = maxNumberParticipants * numberOfSlotsPerLeader;
    }

    // TODO for production purposes it might be desirable to make this function "onlyOwner". However the tests would be
    // harder to write. See ticket https://github.com/EspressoSystems/op-leader-election/issues/73
    function addParticipant(address _addr) public override {
        require(
            indexLastInsertedParticipant < maxNumberParticipants,
            "RoundRobinLeaderElection: list of participants is full."
        );

        participants[indexLastInsertedParticipant] = _addr;
        isParticipant[_addr] = true;
        indexLastInsertedParticipant++;
    }

    function isCurrentLeader(address _leaderId, uint256 _blockNumber) external view override returns (bool) {
        if (_blockNumber < creationBlockNumber) {
            // if the block number is "too old" then no one is leader
            return false;
        } else {
            uint32 numberParticipants = indexLastInsertedParticipant;
            uint256 rangeFromCreation =
                (_blockNumber - creationBlockNumber) % (numberParticipants * numberOfSlotsPerLeader);
            uint256 indexLeader = rangeFromCreation / numberOfSlotsPerLeader;
            return participants[indexLeader] == _leaderId;
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
        if (!this.isParticipant(_leaderId)) {
            bool[] memory emptyArray;
            return (LeaderStatusFlags.Invalid, emptyArray);
        } else {
            // Build the bitmap
            bool[] memory leaderSlots = new bool[](horizon);
            for (uint256 i = 0; i < horizon; i++) {
                leaderSlots[i] = this.isCurrentLeader(_leaderId, _blockNumber + i);
            }

            return (LeaderStatusFlags.Scheduled, leaderSlots);
        }
    }
}
