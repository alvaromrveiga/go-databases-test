1. `rm -rf db/test 2>/dev/null && go test -bench=BenchmarkBadger`
<br>goos: linux
<br>goarch: amd64
<br>pkg: database-benchmarks
<br>BenchmarkBadger/UpdateRead-4                  16          63907479 ns/op
<br>PASS
<br>ok      database-benchmarks     1.058s


2. `CGO_ENABLED=1 rm -rf db/test 2>/dev/null && mkdir db/test && go test -bench=BenchmarkSQLite`
<br>goos: linux
<br>goarch: amd64
<br>pkg: database-benchmarks
<br>BenchmarkSQLite-4              1        13385911692 ns/op
<br>PASS
<br>ok      database-benchmarks     13.393s


<br><br>63907479 ns/op vs 13385911692 ns/op
