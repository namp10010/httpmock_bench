# httpmock_bench
golang benchmarking httpmock and httptest

```bash
  make bench
```

output

```text
goos: darwin
goarch: arm64
pkg: github.com/namp10010/httpmock_bench
Benchmark_httptest-10               4986            335974 ns/op           22864 B/op        188 allocs/op
Benchmark_httpmock-10             813808              1671 ns/op            2384 B/op         25 allocs/op
PASS
ok      github.com/namp10010/httpmock_bench     3.501s
```
