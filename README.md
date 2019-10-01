# grpc-vs-rest


```sh
BenchmarkGRPCWithWokers-12                   50000              22671 ns/op            5059 B/op         88 allocs/op
BenchmarkGRPC-12                        2000000000               0.00 ns/op               0 B/op          0 allocs/op
BenchmarkTlsGRPCWithWokers-12                50000              23601 ns/op            5078 B/op         89 allocs/op
BenchmarkTlsGRPC-12                     2000000000               0.01 ns/op               0 B/op          0 allocs/op
BenchmarkHTTP11GetWithWorker-12              10000            1622764 ns/op            8003 B/op         77 allocs/op
BenchmarkHTTP11Get-12                   2000000000               0.00 ns/op               0 B/op          0 allocs/op
BenchmarkTlsHTTP11GetWithWorker-12            3000            5134027 ns/op           24169 B/op        249 allocs/op
BenchmarkTlsHTTP11Get-12                2000000000               0.01 ns/op               0 B/op          0 allocs/op
```
