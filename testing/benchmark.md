---
layout: default
title: Benchmark
parent: Testing
nav_order: 2
---

# Benchmark

- Test files müssen mit `_test.go` enden:
    ```
    project
    └───example
       │    example.go
       │    example_test.go
    ```
- Tests verwenden das "testing" package
- Methoden Signatur weicht von Unit-Tests ab
    - Kein "Test" Prefix
    - als Argument wird `b *testing.B` verwendet
- für das Ausführen der Tests:
    - `go test -bench=[package]`
    - `go test -bench .`

## Beispiel
- Beispiel Code: [fib benchmark](../examples/benchmark/)
    ```
    $ go test -bench .

    goos: darwin
    goarch: arm64
    pkg: opitz-consulting.com/examples/benchmark
    BenchmarkFib10-10        6436446               180.9 ns/op
    PASS
    ok      opitz-consulting.com/examples/benchmark 1.490s
    ```

## Links
- [benchmarks](https://pkg.go.dev/testing#hdr-Benchmarks)