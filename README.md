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

* No binary dependencies (not even libc)
* No compiler dependencies (no LLVM, no GCC, ...)
* No global state (all mutable variables are local)
* No classes or methods: There is just a) data and b) functions that can operate on data
* No name shadowing, names never change their meaning
* No side effects when importing a package
* No backwards compatibility (we use a rather unique method to ensure everything works)
* Fast compilation (a couple milliseconds should suffice for most programs)
* Small binaries (a "Hello World" program produces a 247-byte binary)
* High performance (should compete with C and Rust)
* Linting embedded into the build process (detects common mistakes and suggests solutions)
* Testing embedded into the language ("q test")
* Formatting tools included ("q format")
* Packages should be highly reusable (not bound to predefined data structures)
* Statically typed with type inference
* User-friendly compiler messages

## Todo

* [x] Tokenizer
* [x] Scanner
* [x] Parallel function compiler
* [x] Error messages
* [x] Function calls
* [x] Infinite loops via `loop`
* [x] Simple `for` loops
* [x] Simple `if` conditions
* [x] `return` with no arguments
* [x] Syscalls
* [x] Linter: Unused variables
* [x] Linter: Unused parameters
* [ ] Function return values
* [ ] Data structures
* [ ] Stack allocation
* [ ] Heap allocation
* [ ] `match` statement
* [ ] Error handling
* [ ] Parallel execution
* [ ] Lock-free data structures
* [ ] Variable lifetime tracking
* [ ] Multiple return values
* [ ] Expression optimization
* [ ] Assembly optimization
* [ ] Linter: Ineffective assignment
* [ ] ...and everything else you'd expect from a modern language.

## Architectures

Currently **Linux x86-64** only. It produces 64-bit ELF binaries.

## Assembly

You can view the produced assembly output via the `-v` verbose flag:

```shell
./q build -v examples/loops
```

Note that this intermediate text representation is only generated when the verbose flag is specified.
Q uses its own binary assembler which avoids the intermediary text format during normal compilation.

## Builtins

There are currently 2 builtin functions, `syscall` and `print`. In the future we'd like to remove `print` so that `syscall` becomes the only builtin function.

## Editors

There is a very simple [VS Code extension](https://github.com/akyoto/vscode-q) that enables syntax highlighting (copy into `$HOME/.vscode/extensions`).

## How to contribute

* TODO.

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
