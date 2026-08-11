# Contributing

To maintain a high-quality codebase, the Q project selectively accepts features.
If you'd like to propose a new feature, please submit a feature proposal that outlines the motivation, design, and expected impact.

## Tests

The test suite covers more than 3,000 test cases.
Passing all tests is the minimum requirement for any proposed change to be considered.

## AI

Every line of code added to this repository must be fully understood by the contributor.
You should be prepared to explain what you're changing, how it affects the rest of the codebase, and what behavioral or performance implications it introduces.
For these reasons, the use of generative AI to produce code contributions is strongly discouraged.

## FAQ

### How do I run the tests?

```shell
# Run all tests
go run gotest.tools/gotestsum@latest

# Generate coverage
go test -coverpkg=./... -coverprofile=cover.out ./...

# View coverage
go tool cover -func cover.out
go tool cover -html cover.out
```

### How do I run the benchmarks?

```shell
# Run compiler benchmarks
go test ./tests -run '^$' -bench . -benchmem

# Run compiler benchmarks in single-threaded mode
GOMAXPROCS=1 go test ./tests -run '^$' -bench . -benchmem

# Generate profiling data
go test ./tests -run '^$' -bench . -benchmem -cpuprofile cpu.out -memprofile mem.out

# View profiling data
go tool pprof -http=:8080 ./cpu.out
go tool pprof -http=:8080 ./mem.out
```

### How do I run a single file in `tests`?

To run a single test file, linter errors must be disabled using the `-no-lint` flag:

```shell
q tests/add.q -no-lint
```

This is needed because tests often assert "obvious" facts that the linter would not allow in normal programs.

### How do I analyze a problem with the compiler?

Replace the `build` command with `ssa` or `asm` to see the intermediate stages which reveal how the compiler understands your program code.
Use `-func` to filter out specific functions.

If that doesn't reveal any bugs, you can also use the excellent [blinkenlights](https://justine.lol/blinkenlights/) from Justine Tunney to step through x86-64 executables one instruction at a time.