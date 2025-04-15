1. rm -rf db/test 2>/dev/null && go test -bench=BenchmarkBadger
goos: linux
goarch: amd64
pkg: database-benchmarks
BenchmarkBadger/UpdateRead-4                  16          63907479 ns/op
PASS
ok      database-benchmarks     1.058s


2. 
```
CGO_ENABLED=1 rm -rf db/test 2>/dev/null && mkdir db/test && go test -bench=BenchmarkSQLite
```
goos: linux
goarch: amd64
pkg: database-benchmarks
BenchmarkSQLite-4              1        13385911692 ns/op
PASS
ok      database-benchmarks     13.393s


63907479 ns/op vs 13385911692 ns/op
