pragma solidity 0.8.19;

import { Test } from "forge-std/Test.sol";

import "../src/L1/RoundRobinLeaderElection.sol";
import "../src/L1/LeaderElectionBatchInbox.sol";

contract RoundRobinLeaderElectionTest is Test {
    RoundRobinLeaderElection leaderContract;
    uint256 public constant N_PARTICIPANTS = 5;
    uint256 DEPLOYMENT_BLOCK_NUMBER = 100;

    function setUp() public {
        vm.roll(DEPLOYMENT_BLOCK_NUMBER);
        leaderContract = new RoundRobinLeaderElection(N_PARTICIPANTS);

        // Poblate the list of participants who are allowed to vote
        for (uint256 i = 1; i <= N_PARTICIPANTS; i++) {
            address addr = vm.addr(i);
            leaderContract.addParticipant(addr);
        }
    }

    function test_addParticipant_participantsAreAdded_success() external {
        for (uint256 i = 0; i < N_PARTICIPANTS; i++) {
            assertEq(vm.addr(i + 1), leaderContract.participants(i));
        }
    }

    function test_addParticipant_listIsFull_reverts() external {
        vm.expectRevert("RoundRobinLeaderElection: list of participants is full.");
        leaderContract.addParticipant(vm.addr(52));
    }

    function test_isLeader_success() external {
        assertEq(block.number, DEPLOYMENT_BLOCK_NUMBER);
        assertEq(leaderContract.creation_block_number(), DEPLOYMENT_BLOCK_NUMBER);

        // No one is a leader for an old block
        for (uint256 i = 1; i <= N_PARTICIPANTS; i++) {
            address participant = leaderContract.participants(i - 1);
            assertFalse(leaderContract.isCurrentLeader(participant, 1));
        }

        // When the contract is deployed the leader is the first participant
        assertTrue(leaderContract.isCurrentLeader(vm.addr(1), DEPLOYMENT_BLOCK_NUMBER));
        assertTrue(leaderContract.isCurrentLeader(vm.addr(2), DEPLOYMENT_BLOCK_NUMBER + 1));
    }

    function test_nextBlocksAsLeader_success() external {
        address notALeader = vm.addr(1234);
        uint256 blockNumber = 0;

        ILeaderElectionBatchInbox.LeaderStatusFlags flag;

        bool[] memory bitmap;

        (flag, bitmap) = leaderContract.nextBlocksAsLeader(notALeader, blockNumber);
        assertTrue(flag == ILeaderElectionBatchInbox.LeaderStatusFlags.Invalid);
        assertTrue(bitmap.length == 0);

        // Happy case for first leader
        blockNumber = DEPLOYMENT_BLOCK_NUMBER;
        address leader = vm.addr(1);

        (flag, bitmap) = leaderContract.nextBlocksAsLeader(leader, blockNumber);
        assertTrue(flag == ILeaderElectionBatchInbox.LeaderStatusFlags.Scheduled);
        assertEq(
            abi.encodePacked(bitmap),
            abi.encodePacked([true, false, false, false, false, true, false, false, false, false])
        );

        // Happy case for third leader
        leader = vm.addr(3);

        assertTrue(flag == ILeaderElectionBatchInbox.LeaderStatusFlags.Scheduled);
        (flag, bitmap) = leaderContract.nextBlocksAsLeader(leader, blockNumber);
        assertEq(
            abi.encodePacked(bitmap),
            abi.encodePacked([false, false, true, false, false, false, false, true, false, false])
        );
    }
}
