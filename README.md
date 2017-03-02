# benchfmt

add `ops/sec` output to go bechmarks:

usage :

```bash
$ go test -bench . | benchfmt 

BenchmarkFoo-4       	   50000	     37698 ns/op	26526 ops/1s
```

or with additional duration param
```bash
$ go test -bench . | benchfmt -d 1h

BenchmarkFoo-4       	   50000	     37807 ns/op	95220461 ops/1h0m0s
```
 
