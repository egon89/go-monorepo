# Learn Go with Tests

[Site](https://quii.gitbook.io/learn-go-with-tests)

```bash
# run all tests
go test -v ./...

# run tests in a specific package
go test ./hello-world

# run benchmarks
go test -bench=.

# run from the root
go test ./... -bench=.

# run benchmarks with memory allocation stats
go test ./... -bench=. -benchmem

# run specific test
go test -run TestAreaTableDrivenTests/Rectangle
```

## Iterate benchmark test
`Strings` in Go are immutable, meaning every concatenation, such as in our Repeat function, involves copying memory to accommodate the new string. This impacts performance, particularly during heavy string concatenation.

The standard library provides the `strings.Builder` type which minimizes memory copying.

![Benchmark Test Image](resources/iteration-benchmark.png)

The `-benchmem` flag reports information about memory allocations:
- B/op: the number of bytes allocated per iteration
- allocs/op: the number of memory allocations per iteration

```bash
go test -bench=.
goos: linux
goarch: amd64
pkg: github.com/egon89/go-monorepo/learn-go-with-tests/concurrency/v2
cpu: ...
BenchmarkCheckWebsites-16             57          20674131 ns/op
PASS
ok      github.com/egon89/go-monorepo/learn-go-with-tests/concurrency/v2        1.202s
```

- goos, goarch, cpu: Your system info (Linux, amd64, CPU model)
- BenchmarkCheckWebsites-16: The name of your benchmark function
  - The -16 means it ran with 16 logical CPUs (parallelism)
- 57: The number of times the benchmark loop ran (b.N)
- 20674131 ns/op: Average time per operation (iteration) in nanoseconds (about 20.67 ms per call)

## Go Coverage tool
The Go toolchain includes a built-in [coverage tool](https://go.dev/blog/cover) that helps you identify which parts of your code are tested by your tests. To use it, you can run:

```bash
go test -cover ./...
```

## var keyword
The var keyword allows us to define values global to the package.

## Anonymous functions
- they can be executed at the same time that they're declared - this is what the `()` at the end of the anonymous function is doing
- all the variables that are available at the point when you declare the anonymous function are also available in the body of the function

## Go Race Detector
Race conditions occur when two or more goroutines access the same variable concurrently, and at least one of the accesses is a write. The Go race detector is a tool that helps identify race conditions in your code. To use it, you can run:

```bash
go test -race ./...
```
```bash
# output
==================
WARNING: DATA RACE
Write at 0x00c000126480 by goroutine 9:
  runtime.mapassign_faststr()
      /home/everton/go/pkg/mod/golang.org/toolchain@v0.0.1-go1.23.4.linux-amd64/src/runtime/map_faststr.go:223 +0x0
  github.com/egon89/go-monorepo/learn-go-with-tests/concurrency.CheckWebsites.func1()
      /home/everton/projects/golang/go-monorepo/learn-go-with-tests/concurrency/v0.go:13 +0x77

Previous write at 0x00c000126480 by goroutine 8:
  runtime.mapassign_faststr()
      /home/everton/go/pkg/mod/golang.org/toolchain@v0.0.1-go1.23.4.linux-amd64/src/runtime/map_faststr.go:223 +0x0
  github.com/egon89/go-monorepo/learn-go-with-tests/concurrency.CheckWebsites.func1()
      /home/everton/projects/golang/go-monorepo/learn-go-with-tests/concurrency/v0.go:13 +0x77
```

In the above output, the race detector has identified a data race in the `CheckWebsites` function, where multiple goroutines are writing to the same memory location concurrently without proper synchronization.

## Channels
- Data structure that can both receive and send values
- Allow communication between different processes
