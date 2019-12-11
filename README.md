# q

[![Godoc][godoc-image]][godoc-url]
[![Report][report-image]][report-url]
[![Tests][tests-image]][tests-url]
[![Coverage][coverage-image]][coverage-url]
[![Sponsor][sponsor-image]][sponsor-url]

This is a **very early version** of a programming language I'm currently working on.

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

* Fast compilation (<1 ms for simple programs)
* Small binaries ("Hello World" is 247 bytes)
* High performance (compete with C and Rust)

### Sub-goals

To achieve these goals, we will implement a new backend with:

* No binary dependencies (not even libc)
* No compiler dependencies (no LLVM, no GCC, ...)

We'll also say goodbye to the following:

* No global state (all mutable variables are local)
* No side effects when importing a package
* No name shadowing, names never change their meaning
* No complicated classes, just simple data structures

Productivity is nice but bugs in your code are not. We'll use:

* Type system (reduce bugs at compile time)
* Linters (reduce bugs at compile time)
* Tests (reduce bugs at test time)
* Contracts (reduce bugs at run time)

Quality of life features:

* Simple dependency management
* Auto-formatter for source code
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
* [x] `import` standard packages
* [x] `expect` for input validation
* [x] `ensure` for output validation
* [ ] Data structures
* [ ] Stack allocation
* [ ] Heap allocation
* [ ] Type system
* [ ] Type operator: `|` (`User | Error`)
* [ ] Hexadecimal, octal and binary literals
* [ ] `match` keyword
* [ ] `import` external packages
* [ ] Error handling
* [ ] Cyclic function calls
* [ ] Multi-threading
* [ ] Lock-free data structures
* [ ] Multiple return values
* [ ] Rewrite compiler in Q
* [ ] ...

### Optimizations

* [x] Exclude unused functions
* [x] Function call inlining
* [x] Assembly optimization backend
* [x] Disable contracts via `-O` flag
* [ ] Expression optimization
* [ ] Loop unrolls
* [ ] ...

### Linter

* [x] Unused variables
* [x] Unused parameters
* [x] Unused imports
* [x] Unmodified mutable variables
* [x] Unnecessary newlines
* [x] Ineffective assignments
* [ ] ...

### Operators

* [x] `+`, `-`, `*`, `/`
* [x] `==`, `!=`, `<`, `<=`, `>`, `>=`
* [x] `=`
* [ ] `+=`, `-=`, `*=`, `/=`
* [ ] `&=`, `|=`
* [ ] `<<=`, `>>=`
* [ ] `<<`, `>>`
* [ ] `&&`, `||`
* [ ] `&`, `|`
* [ ] `%`
* [ ] ...

### Architecture

* [x] x86-64
* [ ] WASM
* [ ] ARM
* [ ] ...

### Platform

* [x] Linux
* [ ] Mac
* [ ] Windows
* [ ] ...

## FAQ

### How do I navigate the source code?

* [benchmarks](https://github.com/akyoto/q/tree/master/benchmarks) contains benchmarks for compilation speed
* [build](https://github.com/akyoto/q/tree/master/build) contains the actual compiler source
* [cli](https://github.com/akyoto/q/tree/master/cli) contains the command line interface
* [examples](https://github.com/akyoto/q/tree/master/examples) contains a few examples that are also used in tests
* [lib](https://github.com/akyoto/q/tree/master/lib) contains the standard library

### How do I view the produced assembly output?

```shell
q build -a
q build --assembly
```

### How can I make a performance optimized build?

```shell
q build -O
q build --optimize
```

This will disable all `expect` and `ensure` checks.

### How can I see where my compilation time is spent on?

```shell
q build -t
q build --time
```

### How do I install it system-wide?

```shell
sudo ln -s $PWD/q /usr/local/bin/q
```

### How can I include information about my system in bug reports?

```shell
q system
```

### Which editor can I use to edit Q code?

There is a simple [VS Code extension](https://github.com/akyoto/vscode-q) with syntax highlighting.

```shell
git clone https://github.com/akyoto/vscode-q ~/.vscode/extensions/vscode-q
```

### Is the syntax final?

Unlikely. There will be changes in the near future.

### Which builtin functions are available?

There are currently 2 builtin functions, `syscall` and `print`. In the future we'd like to remove `print` so that `syscall` becomes the only builtin function.

### How do I run the tests?

```shell
go test -coverpkg=./...
```

### How do I run the benchmarks?

```shell
go test -bench=. ./benchmarks
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
