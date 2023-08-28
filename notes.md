# Notes

Update to SystemConfig happens in two places

- l1_traversal
- building of payload attributes

- For the l1_traversal the the `sysCfg` attribute is modified.
- For payload attributes the result of applying the update does not seem to be persistet.

calldata: 01234567890123456789012345678901234567890123456789

    [PASS] test_fallback_batch_success() (gas: 21471)
    Traces:
    [21471] RoundRobinLeaderElectionTest::test_fallback_batch_success()
        ├─ [0] VM::addr(<pk>) [staticcall]
        │   └─ ← 0x7E5F4552091A69125d5DfCb7b8C2659029395Bdf
        ├─ [0] VM::prank(0x7E5F4552091A69125d5DfCb7b8C2659029395Bdf)
        │   └─ ← ()
        ├─ [13170] Proxy::30313233(34353637383930313233343536373839303132333435363738393031323334353637383930313233343536373839)
        │   ├─ [8208] RoundRobinLeaderElection.0.8.15::30313233(34353637383930313233343536373839303132333435363738393031323334353637383930313233343536373839) [delegatecall]
        │   │   ├─ [7599] Proxy::isCurrentLeader(0x7E5F4552091A69125d5DfCb7b8C2659029395Bdf, 100) [staticcall]
        │   │   │   ├─ [7110] RoundRobinLeaderElection.0.8.15::isCurrentLeader(0x7E5F4552091A69125d5DfCb7b8C2659029395Bdf, 100) [delegatecall]
        │   │   │   │   └─ ← true
        │   │   │   └─ ← true
        │   │   └─ ← ()
        │   └─ ← ()
        └─ ← ()

    [PASS] test_submit_batch_success() (gas: 21918)
    Traces:
    [21918] RoundRobinLeaderElectionTest::test_submit_batch_success()
        ├─ [0] VM::addr(<pk>) [staticcall]
        │   └─ ← 0x7E5F4552091A69125d5DfCb7b8C2659029395Bdf
        ├─ [0] VM::prank(0x7E5F4552091A69125d5DfCb7b8C2659029395Bdf)
        │   └─ ← ()
        ├─ [13489] Proxy::submit(0x3031323334353637383930313233343536373839303132333435363738393031323334353637383930313233343536373839) [staticcall]
        │   ├─ [8491] RoundRobinLeaderElection.0.8.15::submit(0x3031323334353637383930313233343536373839303132333435363738393031323334353637383930313233343536373839) [delegatecall]
        │   │   ├─ [7599] Proxy::isCurrentLeader(0x7E5F4552091A69125d5DfCb7b8C2659029395Bdf, 100) [staticcall]
        │   │   │   ├─ [7110] RoundRobinLeaderElection.0.8.15::isCurrentLeader(0x7E5F4552091A69125d5DfCb7b8C2659029395Bdf, 100) [delegatecall]
        │   │   │   │   └─ ← true
        │   │   │   └─ ← true
        │   │   └─ ← ()
        │   └─ ← ()
        └─ ← ()

calldata: 0123456789

    [PASS] test_fallback_batch_success() (gas: 21453)
    Traces:
    [21453] RoundRobinLeaderElectionTest::test_fallback_batch_success()
        ├─ [0] VM::addr(<pk>) [staticcall]
        │   └─ ← 0x7E5F4552091A69125d5DfCb7b8C2659029395Bdf
        ├─ [0] VM::prank(0x7E5F4552091A69125d5DfCb7b8C2659029395Bdf)
        │   └─ ← ()
        ├─ [13167] Proxy::30313233(343536373839)
        │   ├─ [8208] RoundRobinLeaderElection.0.8.15::30313233(343536373839) [delegatecall]
        │   │   ├─ [7599] Proxy::isCurrentLeader(0x7E5F4552091A69125d5DfCb7b8C2659029395Bdf, 100) [staticcall]
        │   │   │   ├─ [7110] RoundRobinLeaderElection.0.8.15::isCurrentLeader(0x7E5F4552091A69125d5DfCb7b8C2659029395Bdf, 100) [delegatecall]
        │   │   │   │   └─ ← true
        │   │   │   └─ ← true
        │   │   └─ ← ()
        │   └─ ← ()
        └─ ← ()

    [PASS] test_submit_batch_success() (gas: 21894)
    Traces:
    [21894] RoundRobinLeaderElectionTest::test_submit_batch_success()
        ├─ [0] VM::addr(<pk>) [staticcall]
        │   └─ ← 0x7E5F4552091A69125d5DfCb7b8C2659029395Bdf
        ├─ [0] VM::prank(0x7E5F4552091A69125d5DfCb7b8C2659029395Bdf)
        │   └─ ← ()
        ├─ [13483] Proxy::submit(0x30313233343536373839) [staticcall]
        │   ├─ [8491] RoundRobinLeaderElection.0.8.15::submit(0x30313233343536373839) [delegatecall]
        │   │   ├─ [7599] Proxy::isCurrentLeader(0x7E5F4552091A69125d5DfCb7b8C2659029395Bdf, 100) [staticcall]
        │   │   │   ├─ [7110] RoundRobinLeaderElection.0.8.15::isCurrentLeader(0x7E5F4552091A69125d5DfCb7b8C2659029395Bdf, 100) [delegatecall]
        │   │   │   │   └─ ← true
        │   │   │   └─ ← true
        │   │   └─ ← ()
        │   └─ ← ()
        └─ ← ()


calldata: empty

    [PASS] test_fallback_batch_success() (gas: 21227)
    Traces:
    [21227] RoundRobinLeaderElectionTest::test_fallback_batch_success()
        ├─ [0] VM::addr(<pk>) [staticcall]
        │   └─ ← 0x7E5F4552091A69125d5DfCb7b8C2659029395Bdf
        ├─ [0] VM::prank(0x7E5F4552091A69125d5DfCb7b8C2659029395Bdf)
        │   └─ ← ()
        ├─ [12965] Proxy::receive()
        │   ├─ [8084] RoundRobinLeaderElection.0.8.15::fallback() [delegatecall]
        │   │   ├─ [7599] Proxy::isCurrentLeader(0x7E5F4552091A69125d5DfCb7b8C2659029395Bdf, 100) [staticcall]
        │   │   │   ├─ [7110] RoundRobinLeaderElection.0.8.15::isCurrentLeader(0x7E5F4552091A69125d5DfCb7b8C2659029395Bdf, 100) [delegatecall]
        │   │   │   │   └─ ← true
        │   │   │   └─ ← true
        │   │   └─ ← ()
        │   └─ ← ()
        └─ ← ()

    [PASS] test_submit_batch_success() (gas: 21870)
    Traces:
    [21870] RoundRobinLeaderElectionTest::test_submit_batch_success()
        ├─ [0] VM::addr(<pk>) [staticcall]
        │   └─ ← 0x7E5F4552091A69125d5DfCb7b8C2659029395Bdf
        ├─ [0] VM::prank(0x7E5F4552091A69125d5DfCb7b8C2659029395Bdf)
        │   └─ ← ()
        ├─ [13477] Proxy::submit(0x) [staticcall]
        │   ├─ [8491] RoundRobinLeaderElection.0.8.15::submit(0x) [delegatecall]
        │   │   ├─ [7599] Proxy::isCurrentLeader(0x7E5F4552091A69125d5DfCb7b8C2659029395Bdf, 100) [staticcall]
        │   │   │   ├─ [7110] RoundRobinLeaderElection.0.8.15::isCurrentLeader(0x7E5F4552091A69125d5DfCb7b8C2659029395Bdf, 100) [delegatecall]
        │   │   │   │   └─ ← true
        │   │   │   └─ ← true
        │   │   └─ ← ()
        │   └─ ← ()
        └─ ← ()
