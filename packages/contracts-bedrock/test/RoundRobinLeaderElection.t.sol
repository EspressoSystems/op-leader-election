pragma solidity 0.8.15;

import {Test} from "forge-std/Test.sol";

import "../src/L1/RoundRobinLeaderElection.sol";

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

    function test_rrElection_ListParticipants() external {
        for (uint256 i = 0; i < N_PARTICIPANTS; i++) {
            assertEq(vm.addr(i + 1), leaderContract.participants(i));
        }
    }

    function test_rrElection_ListParticipantsIsFull() external {
        vm.expectRevert("List of participants is full.");
        leaderContract.addParticipant(vm.addr(52));
    }

    function test_rrElection_IsLeader() external {
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
}
