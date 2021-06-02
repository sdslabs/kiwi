![Kiwi Logo](./docs/.vuepress/public/kiwi-logo.png)

> A minimalistic in-memory key value store.

![Go CI](https://github.com/sdslabs/kiwi/workflows/Go%20CI/badge.svg) ![Docs CI](https://github.com/sdslabs/kiwi/workflows/Docs%20CI/badge.svg) ![Docs CD](https://github.com/sdslabs/kiwi/workflows/Docs%20CD/badge.svg) [![PkgGoDev](https://pkg.go.dev/badge/github.com/sdslabs/kiwi)](https://pkg.go.dev/github.com/sdslabs/kiwi)

## Overview

You can think of Kiwi as thread safe global variables. This kind of library
comes in helpful when you need to manage state across your application which
can be mutated with multiple threads. Kiwi protects your keys with mutex locks
so you don't have to.

Head over to [kiwi.sdslabs.co](https://kiwi.sdslabs.co) for more details and
documentation.

## Installation

> Kiwi requires [Go](https://golang.org/) >= 1.14

Kiwi can be integrated with your application just like any other go library.

```sh
go get -u github.com/sdslabs/kiwi
```

Now you can import kiwi any where in your code.

```go
import "github.com/sdslabs/kiwi"
```

## Basic usage

Create a store, add key and play with it. It's that easy!

```go
store := stdkiwi.NewStore()

if err := store.AddKey("my_string", "str"); err != nil {
  // handle error
}

myString := store.Str("my_string")

if err := myString.Update("Hello, World!"); err != nil {
  // handle error
}

str, err := myString.Get()
if err != nil {
  // handle error
}

fmt.Println(str) // Hello, World!
```

Check out the [tutorial](https://kiwi.sdslabs.co/docs/tutorial-store.html) to
learn how to use Kiwi.

## Benchmarks

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

## Contributing

We are always open for contributions. If you find any feature missing, or just
want to report a bug, feel free to open an issue and/or submit a pull request
regarding the same.

For more information on contribution, check out our
[docs](https://kiwi.sdslabs.co/docs/contribution-guide.html).

## Contact

If you have a query regarding the product or just want to say hello then feel
free to visit [chat.sdslabs.co](https://chat.sdslabs.co) or drop a mail at
[contact@sdslabs.co.in](mailto:contact@sdslabs.co.in)

---

Made by [SDSLabs](https://sdslabs.co)
