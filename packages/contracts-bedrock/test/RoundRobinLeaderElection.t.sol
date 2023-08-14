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

    // TODO is there a better way to compare these arrays?
    function compareArrays(uint8[] memory _first_arr, uint8[10] memory _second_arr) private pure returns (bool) {
        for (uint256 i = 0; i < 10; i++) {
            if (_first_arr[i] != _second_arr[i]) {
                return false;
            }
        }
        return true;
    }

    function test_nextBlocksAsLeader_success() external {
        // If the caller is not as participant return the tuple (LeaderStatus.Invalid, 0,0,0)
        address notALeader = vm.addr(1234);
        vm.prank(notALeader);

        ILeaderElectionBatchInbox.LeaderStatusFlags flag;
        uint256 blockNumber;
        uint8[] memory bitmap;

        (flag, blockNumber, bitmap) = leaderContract.nextBlocksAsLeader();
        assertTrue(flag == ILeaderElectionBatchInbox.LeaderStatusFlags.Invalid);
        assertEq(0, blockNumber);
        assertTrue(bitmap.length == 0);

        // Happy case for first leader
        address leader = vm.addr(1);
        vm.prank(leader);

        (flag, blockNumber, bitmap) = leaderContract.nextBlocksAsLeader();
        assertTrue(flag == ILeaderElectionBatchInbox.LeaderStatusFlags.Scheduled);
        assertEq(100, blockNumber);
        assertTrue(compareArrays(bitmap, [1, 0, 0, 0, 0, 1, 0, 0, 0, 0]));

        // Happy case for third leader
        leader = vm.addr(3);
        vm.prank(leader);

        assertTrue(flag == ILeaderElectionBatchInbox.LeaderStatusFlags.Scheduled);
        (flag, blockNumber, bitmap) = leaderContract.nextBlocksAsLeader();
        assertEq(100, blockNumber);
        assertTrue(compareArrays(bitmap, [0, 0, 1, 0, 0, 0, 0, 1, 0, 0]));
    }
}
