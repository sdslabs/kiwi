# Benchmarks

Following are the benchmarks comparing Kiwi with BuntDB on a MacBook Pro (8th gen
Intel i5 2.4GHz processor, 8GB RAM).

```
❯ go test -bench=. -test.benchmem ./benchmark
goos: darwin
goarch: amd64
pkg: github.com/sdslabs/kiwi/benchmark
BenchmarkBuntDB_Update-8        11777931                96.6 ns/op            48 B/op          1 allocs/op
BenchmarkBuntDB_View-8          23310963                47.1 ns/op            48 B/op          1 allocs/op
BenchmarkKiwi_Update-8          10356004               115 ns/op              48 B/op          3 allocs/op
BenchmarkKiwi_Get-8             21910110                53.2 ns/op             0 B/op          0 allocs/op
PASS
ok      github.com/sdslabs/kiwi/benchmark       6.216s
```

Following are the key differences due to which Kiwi is a little slow:

1. BuntDB supports transactions, i.e., it locks the database once to apply all
   the operations (and this is what is tested).
1. Kiwi supports dynamic data-types, which means, allocation on heap at runtime
   (`interface{}`) whereas BuntDB is statically typed.
   
The above two differences are what makes Kiwi unique and suitable to use on
many occasions. Due to the aforementioned reasons, Kiwi can support typed values
and not everything is just another "string".

There are places where we could improve more. Some performance issues also lie
in the implementation of values. For example, when updating a string, not returning
the updated string avoids an extra allocation.
