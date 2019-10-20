# grpc-vs-rest


```sh
 grpc-vs-rest git:(master) ✗ go test -run=^$ -bench=HTTP2 -benchtime=10s
goos: darwin
goarch: amd64
pkg: github.com/jbactad/grpc-vs-rest
BenchmarkGRPCHTTP2_4-12         100000            113240 ns/op
BenchmarkGRPCHTTP2_8-12         200000             87792 ns/op
BenchmarkGRPCHTTP2_16-12        200000             67132 ns/op
BenchmarkGRPCHTTP2_32-12        200000             51553 ns/op
BenchmarkGRPCHTTP2_64-12        300000             48865 ns/op
BenchmarkGRPCHTTP2_128-12       500000             46179 ns/op
BenchmarkGRPCHTTP2_256-12       200000             50788 ns/op
BenchmarkGRPCHTTP2_512-12       200000             51909 ns/op
BenchmarkRestHTTP2_4-12         100000            166760 ns/op
BenchmarkRestHTTP2_8-12         100000            138141 ns/op
BenchmarkRestHTTP2_16-12        100000            119753 ns/op
BenchmarkRestHTTP2_32-12        200000            105505 ns/op
BenchmarkRestHTTP2_64-12        200000             98870 ns/op
BenchmarkRestHTTP2_128-12       200000             93604 ns/op
BenchmarkRestHTTP2_256-12       200000             93770 ns/op
BenchmarkRestHTTP2_512-12       200000             93254 ns/op
PASS
ok    github.com/jbactad/grpc-vs-rest 267.026s
```

