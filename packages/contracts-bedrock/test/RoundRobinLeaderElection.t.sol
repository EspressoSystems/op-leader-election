pragma solidity ^0.8.0;

import { Test } from "forge-std/Test.sol";

import { Proxy } from "../src/universal/Proxy.sol";
import { RoundRobinLeaderElection } from "../src/L1/RoundRobinLeaderElection.sol";
import { LeaderElectionBatchInbox } from "../src/L1/LeaderElectionBatchInbox.sol";

contract RoundRobinLeaderElectionTest is Test {
    RoundRobinLeaderElection leaderContract;
    uint256 public constant N_PARTICIPANTS = 4;
    uint256 public constant N_SLOTS_PER_LEADER = 3;
    uint256 DEPLOYMENT_BLOCK_NUMBER = 100;
    address owner = address(0xbeef);

    function setUp() public {
        vm.roll(DEPLOYMENT_BLOCK_NUMBER);

        Proxy proxy = new Proxy(msg.sender);
        RoundRobinLeaderElection leaderImpl = new RoundRobinLeaderElection();

        vm.prank(msg.sender);
        proxy.upgradeToAndCall(
            address(leaderImpl), abi.encodeCall(leaderImpl.initialize, (owner, N_PARTICIPANTS, N_SLOTS_PER_LEADER))
        );

        leaderContract = RoundRobinLeaderElection(address(proxy));

        // Populate the list of participants who are allowed to vote
        for (uint256 i = 1; i <= N_PARTICIPANTS; i++) {
            vm.prank(owner);
            address addr = vm.addr(i);
            leaderContract.addParticipant(addr);
        }
        assertEq(N_PARTICIPANTS * N_SLOTS_PER_LEADER, leaderContract.horizon());
    }

    function test_addParticipant_participantsAreAdded_success() external {
        for (uint256 i = 0; i < N_PARTICIPANTS; i++) {
            assertEq(vm.addr(i + 1), leaderContract.participants(i));
        }
    }

    function test_addParticipant_listIsFull_reverts() external {
        vm.expectRevert("RoundRobinLeaderElection: list of participants is full.");
        vm.prank(owner);
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
        for (uint256 i = 0; i < N_PARTICIPANTS; i++) {
            for (uint256 j = 0; j < N_SLOTS_PER_LEADER; j++) {
                uint256 blockNumber = DEPLOYMENT_BLOCK_NUMBER + i * N_SLOTS_PER_LEADER + j;
                assertTrue(leaderContract.isCurrentLeader(vm.addr(i + 1), blockNumber));
                blockNumber = DEPLOYMENT_BLOCK_NUMBER + 3 * leaderContract.horizon() + i * N_SLOTS_PER_LEADER + j;
                assertTrue(leaderContract.isCurrentLeader(vm.addr(i + 1), blockNumber));
            }
        }
    }

    function test_submit_batch_success() external {
        // Wrong leader, cannot submit
        address notALeader = vm.addr(1234);
        vm.prank(notALeader);
        vm.expectRevert("LeaderElectionBatchInbox: submit function must be called by the leader.");

        LeaderElectionBatchInbox.Meta[] memory metas = new LeaderElectionBatchInbox.Meta[](1);
        leaderContract.submit(metas, "frames");

        // Correct leader, does not revert
        vm.prank(vm.addr(1));
        leaderContract.submit(metas, "frames");
    }

    function test_nextBlocksAsLeader_success() external {
        address notALeader = vm.addr(1234);
        uint256 blockNumber = 0;

        LeaderElectionBatchInbox.LeaderStatusFlags flag;

        bool[] memory bitmap;

        (flag, bitmap) = leaderContract.nextBlocksAsLeader(notALeader, blockNumber);
        assertTrue(flag == LeaderElectionBatchInbox.LeaderStatusFlags.Invalid);
        assertTrue(bitmap.length == 0);

        // Happy case for first leader
        blockNumber = DEPLOYMENT_BLOCK_NUMBER;
        address leader = vm.addr(1);
        assertTrue(leaderContract.is_participant(leader));

        (flag, bitmap) = leaderContract.nextBlocksAsLeader(leader, blockNumber);
        assertTrue(flag == LeaderElectionBatchInbox.LeaderStatusFlags.Scheduled);

        assertEq(bitmap.length, leaderContract.horizon());

        assertEq(
            abi.encodePacked(bitmap),
            abi.encodePacked([true, true, true, false, false, false, false, false, false, false, false, false])
        );

        // Happy case for third leader
        leader = vm.addr(3);
        blockNumber = DEPLOYMENT_BLOCK_NUMBER + 3;
        assertTrue(flag == LeaderElectionBatchInbox.LeaderStatusFlags.Scheduled);
        (flag, bitmap) = leaderContract.nextBlocksAsLeader(leader, blockNumber);
        assertEq(
            abi.encodePacked(bitmap),
            abi.encodePacked([false, false, false, true, true, true, false, false, false, false, false, false])
        );
    }
}
