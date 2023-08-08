pragma solidity 0.8.15;

import { Test } from "forge-std/Test.sol";

import "../src/L1/RoundRobinLeaderElection.sol";

contract RoundRobinLeaderElectionTest is Test {

    RoundRobinLeaderElection leaderContract;

    function setUp() public {
        leaderContract = new RoundRobinLeaderElection();
    }

    function testIsLeader() external {
        assertTrue(leaderContract.isCurrentLeader(0x0000000000000000000000000000000000000000,0));
    }

}
