## Go test -bench
1. `rm -rf db/test 2>/dev/null && go test -bench=BenchmarkBadger`
<br>goos: linux
<br>goarch: amd64
<br>pkg: database-benchmarks
<br>BenchmarkBadger/UpdateRead-4                  16          **63907479 ns/op**
<br>PASS
<br>ok      database-benchmarks     **1.058s**

1. `CGO_ENABLED=1 rm -rf db/test 2>/dev/null && mkdir db/test && go test -bench=BenchmarkSQLite`
##### Parallel:
goos: linux
<br>goarch: amd64
<br>pkg: database-benchmarks
<br>BenchmarkSQLite-4              1        **13385911692 ns/op**
<br>PASS
<br>ok      database-benchmarks     **13.393s**

##### Sequential:
goos: linux
<br>goarch: amd64
<br>pkg: database-benchmarks
<br>cpu: Intel(R) Core(TM) i5-6200U CPU @ 2.30GHz
<br>BenchmarkSQLite/UpdateRead-4                   1        **4006846970 ns/op**
<br>PASS
<br>ok      database-benchmarks     **12.910s**

**63907479 ns/op vs 13385911692 ns/op**
<br><u>**1.058s vs 12.910s**</u>

## Go run
3. `CGO_ENABLED=1 go run .`

##### Output without logging reads:
2025/04/15 15:40:00 Starting BadgerDB benchmark
<br>badger 2025/04/15 15:40:00 INFO: All 0 tables opened in 0s
<br>badger 2025/04/15 15:40:00 INFO: Discard stats nextEmptySlot: 0
<br>badger 2025/04/15 15:40:00 INFO: Set nextTxnTs to 0
<br>2025/04/15 15:40:00 **Badger total time: 291.941044ms**
<br>2025/04/15 15:40:00 Starting SQLite3 benchmark
<br>2025/04/15 15:40:32 SQLite error 0: No user found with id 68
<br>2025/04/15 15:40:53 **SQLite3 total time:  52.641854762s**
<br>badger 2025/04/15 15:40:53 INFO: Lifetime L0 stalled for: 0s
<br>badger 2025/04/15 15:40:53 INFO:
<br>Level 0 [ ]: NumTables: 01. Size: 216 KiB of 0 B. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 64 MiB
<br>Level 1 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
<br>Level 2 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
<br>Level 3 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
<br>Level 4 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
<br>Level 5 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
<br>Level 6 [B]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
<br>Level Done

<u>**291.941044ms vs 52.641854762s**</u>

##### Output with logging reads (read logs are omitted for brevity):
2025/04/15 15:43:44 Starting BadgerDB benchmark
<br>badger 2025/04/15 15:43:44 INFO: All 0 tables opened in 0s
<br>badger 2025/04/15 15:43:44 INFO: Discard stats nextEmptySlot: 0
<br>badger 2025/04/15 15:43:44 INFO: Set nextTxnTs to 0
<br>2025/04/15 15:43:45 **Badger total time: 589.987159ms**
<br>2025/04/15 15:43:45 Starting SQLite3 benchmark
<br>2025/04/15 15:44:46 **SQLite3 total time:  1m0.951128793s**
<br>badger 2025/04/15 15:44:46 INFO: Lifetime L0 stalled for: 0s
<br>badger 2025/04/15 15:44:46 INFO:
<br>Level 0 [ ]: NumTables: 01. Size: 217 KiB of 0 B. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 64 MiB
<br>Level 1 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
<br>Level 2 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
<br>Level 3 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
<br>Level 4 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
<br>Level 5 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
<br>Level 6 [B]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
<br>Level Done

<u>**589.987159ms vs 1m0.951128793s**</u>
