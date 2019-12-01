# q

[![Godoc][godoc-image]][godoc-url]
[![Report][report-image]][report-url]
[![Tests][tests-image]][tests-url]
[![Coverage][coverage-image]][coverage-url]
[![Sponsor][sponsor-image]][sponsor-url]

This is a **very early version** of a programming language I'm currently working on.

It produces machine code and ELF binaries directly without using an external assembler or compiler.

## Installation

```shell
git clone https://github.com/akyoto/q
cd q
go build
```

This will produce the `q` compiler in your current directory.

```shell
./q build examples/hello
./examples/hello/hello
```

## Goals

### No...

* No binary dependencies (not even libc)
* No compiler dependencies (no LLVM, no GCC, ...)
* No global state (all mutable variables are local)
* No classes or methods: There is just a) data and b) functions that can operate on data
* No name shadowing, names never change their meaning
* No side effects when importing a package

### Yes...

* Fast compilation (less than 1 millisecond for simple programs)
* Small binaries ("Hello World" is 247 bytes)
* High performance (compete with C and Rust)
* Linting (detects common mistakes)
* Testing ("q test")
* Formatting ("q fmt")
* Contracts (Eiffel-style)
* Packages should be highly reusable
* Statically typed with type inference
* User-friendly compiler messages

## Todo

### Compiler

* [x] Tokenizer
* [x] Scanner
* [x] Parallel function compiler
* [x] Error messages
* [x] Expression parser
* [x] Function calls
* [x] Infinite `loop`
* [x] Simple `for` loops
* [x] Simple `if` conditions
* [x] Syscalls
* [x] Detect pure functions
* [x] Immutable variables
* [x] Mutable variables via `mut`
* [x] Variable lifetime tracking
* [x] `return` values
* [ ] `import` packages
* [ ] Data structures
* [ ] Stack allocation
* [ ] Heap allocation
* [ ] Cyclic function calls
* [ ] `require` for input validation
* [ ] `ensure` for output validation
* [ ] `match` keyword
* [ ] Error handling
* [ ] Parallelism
* [ ] Lock-free data structures
* [ ] Multiple return values
* [ ] Expression optimization
* [ ] Assembly optimization
* [ ] ...

### Linter

* [x] Unused variables
* [x] Unused parameters
* [ ] Ineffective assignment
* [ ] ...

### Architecture

* [x] x86-64
* [ ] WebAssembly
* [ ] ARM
* [ ] ...

### OS

* [x] Linux
* [ ] Mac
* [ ] Windows
* [ ] ...

## FAQ

### How do I navigate the source code?

* [benchmarks](https://github.com/akyoto/q/tree/master/benchmarks) contains tests for compilation speed
* [build](https://github.com/akyoto/q/tree/master/build) contains the actual compiler source
* [cli](https://github.com/akyoto/q/tree/master/cli) contains the command line interface
* [examples](https://github.com/akyoto/q/tree/master/examples) contains a few examples that are also used in tests
* [lib](https://github.com/akyoto/q/tree/master/lib) contains the standard library

### How do I run the tests?

```shell
go test -v ./...
```

### How do I run the benchmarks?

```shell
go test -run=^$ -bench=. ./...
```

### Can I view the produced assembly output?

Yes, use the `-v` verbose flag:

```shell
q build -v examples/loops
```

### Is the syntax final?

No, the syntax will be changed in the future.

### Which builtin functions are available?

There are currently 2 builtin functions, `syscall` and `print`. In the future we'd like to remove `print` so that `syscall` becomes the only builtin function.

### Which editor can I use to edit Q code?

There is a simple [VS Code extension](https://github.com/akyoto/vscode-q) with syntax highlighting.

```shell
git clone https://github.com/akyoto/vscode-q ~/.vscode/extensions/vscode-q
```

### Is there a community for this project?

There is a Discord channel and a Telegram group for [sponsors](https://github.com/sponsors/akyoto).

## Style

Please take a look at the [style guidelines](https://github.com/akyoto/quality/blob/master/STYLE.md) if you'd like to make a pull request.

## Sponsors

| [![Cedric Fung](https://avatars3.githubusercontent.com/u/2269238?s=70&v=4)](https://github.com/cedricfung) | [![Scott Rayapoullé](https://avatars3.githubusercontent.com/u/11772084?s=70&v=4)](https://github.com/soulcramer) | [![Eduard Urbach](https://avatars3.githubusercontent.com/u/438936?s=70&v=4)](https://eduardurbach.com) |
| --- | --- | --- |
| [Cedric Fung](https://github.com/cedricfung) | [Scott Rayapoullé](https://github.com/soulcramer) | [Eduard Urbach](https://eduardurbach.com) |

Want to see [your own name here?](https://github.com/users/akyoto/sponsorship)

[godoc-image]: https://godoc.org/github.com/akyoto/q?status.svg
[godoc-url]: https://godoc.org/github.com/akyoto/q
[report-image]: https://goreportcard.com/badge/github.com/akyoto/q
[report-url]: https://goreportcard.com/report/github.com/akyoto/q
[tests-image]: https://cloud.drone.io/api/badges/akyoto/q/status.svg
[tests-url]: https://cloud.drone.io/akyoto/q
[coverage-image]: https://codecov.io/gh/akyoto/q/graph/badge.svg
[coverage-url]: https://codecov.io/gh/akyoto/q
[sponsor-image]: https://img.shields.io/badge/github-donate-green.svg
[sponsor-url]: https://github.com/users/akyoto/sponsorship
