<div align="center">
	<img src="logo.svg" width="90" alt="q logo">
	<h1>The Q Programming Language</h1>
	<p>Q is a minimal, dependency-free programming language and compiler targeting x86-64 and arm64 with ultra-fast builds and tiny binaries.</p>
</div>

## Features

* High performance (`ssa` and `asm` optimizations)
* Fast compilation (<100 μs for simple programs)
* Tiny executables ("Hello World" is ~600 bytes)
* Multiple platforms (Linux, Mac and Windows)
* Zero dependencies (no llvm, no libc)

## Installation

> [!WARNING]
> Q is under heavy development and not ready for production yet.
>
> Please read the [comment on the status](https://lobste.rs/s/t7osqo/q_programming_language) of the project.
>
> Feel free to [contact me](https://urbach.dev/contact) if you are interested in helping out.

Build from source:

```shell
git clone https://git.urbach.dev/cli/q
cd q
go build
```

Install via symlink:

```shell
ln -s $PWD/q ~/.local/bin/q
```

## Usage

Run:

```shell
q examples/hello
```

Build:

```shell
q build examples/hello
```

Cross-compile:

```shell
q build examples/hello --os [linux|mac|windows] --arch [x86|arm]
```

## News

- **2025-08-19**: Performance improvements.
- **2025-08-18**: Slices for strings.
- **2025-08-17**: Struct allocation by value/reference.
- **2025-08-16**: Multiple return values.
- **2025-08-15**: Data structures.
- **2025-08-14**: Memory load and store instructions.
- **2025-08-13**: Naive memory allocations.
- **2025-08-12**: Support for Windows on arm64.
- **2025-08-11**: Support for Mac on arm64.

## Examples

The syntax is still highly unstable because I'm focusing my work on the correct machine code generation for all platforms and architectures. However, you can take a look at the [examples](../examples) and the [tests](../tests) to get a perspective on the current status.

- [hello](../examples/hello/hello.q)
- [echo](../examples/echo/echo.q)
- [fibonacci](../examples/fibonacci/fibonacci.q)
- [fizzbuzz](../examples/fizzbuzz/fizzbuzz.q)

## Cheat Sheet

| I need to...                     |                             | API stability   |
| -------------------------------- | --------------------------- | --------------- |
| Define a new variable            | `x := 1`                    | ✔️ Stable       |
| Reassign an existing variable    | `x = 2`                     | ✔️ Stable       |
| Define a function                | `main() {}`                 | ✔️ Stable       |
| Define a struct                  | `Point {}`                  | ✔️ Stable       |
| Define input and output types    | `f(a int) -> (b int) {}`    | ✔️ Stable       |
| Instantiate a struct             | `Point{x: 1, y: 2}`         | ✔️ Stable       |
| Instantiate a struct on the heap | `new(Point)`                | 🚧 Experimental |
| Access struct fields             | `p.x`                       | ✔️ Stable       |
| Dereference a pointer            | `[ptr]`                     | ✔️ Stable       |
| Index a pointer                  | `ptr[0]`                    | ✔️ Stable       |
| Slice a string                   | `"Hello"[1..3]`             | ✔️ Stable       |
| Return multiple values           | `return 1, 2`               | ✔️ Stable       |
| Loop                             | `loop {}`                   | ✔️ Stable       |
| Loop 10 times                    | `loop 0..10 {}`             | ✔️ Stable       |
| Loop 10 times with a variable    | `loop i := 0..10 {}`        | ✔️ Stable       |
| Branch multiple conditions       | `switch{ ... }`             | ✔️ Stable       |
| Define a constant                | `const { x 42 }`            | 🚧 Experimental |
| Define an extern C function      | `extern { g { f() } }`      | ✔️ Stable       |
| Allocate memory                  | `mem.alloc(4096)`           | ✔️ Stable       |
| Free memory                      | `mem.free(buffer)`          | 🚧 Experimental |
| Output a string                  | `io.write("Hello\n")`       | ✔️ Stable       |
| Output an integer                | `io.writeInt(42)`           | 🚧 Experimental |

## Source

The source code structure uses a flat layout without nesting:

- [arm](../src/arm) - arm64 architecture
- [asm](../src/asm) - Generic assembler
- [ast](../src/ast) - Abstract syntax tree
- [cli](../src/cli) - Command line interface
- [codegen](../src/codegen) - SSA to assembly code generation
- [compiler](../src/compiler) - Compiler frontend
- [config](../src/config) - Build configuration
- [core](../src/core) - Defines `Function` and compiles tokens to SSA
- [cpu](../src/cpu) - Types to represent a generic CPU
- [data](../src/data) - Data container that can re-use existing data
- [dll](../src/dll) - DLL support for Windows systems
- [elf](../src/elf) - ELF format for Linux executables
- [errors](../src/errors) - Error handling that reports lines and columns
- [exe](../src/exe) - Generic executable format to calculate section offsets
- [expression](../src/expression) - Expression parser generating trees
- [fs](../src/fs) - File system access
- [global](../src/global) - Global variables like the working directory
- [linker](../src/linker) - Frontend for generating executable files
- [macho](../src/macho) - Mach-O format for Mac executables
- [memfile](../src/memfile) - Memory backed file descriptors
- [pe](../src/pe) - PE format for Windows executables
- [scanner](../src/scanner) - Scanner that parses top-level instructions
- [set](../src/set) - Generic set implementation
- [sizeof](../src/sizeof) - Calculates the byte size of numbers
- [ssa](../src/ssa) - Static single assignment types
- [token](../src/token) - Tokenizer
- [types](../src/types) - Type system
- [verbose](../src/verbose) - Verbose output
- [x86](../src/x86) - x86-64 architecture

The typical flow for a build command is the following:

1. [main](../main.go)
1. [cli.Exec](../src/cli/Exec.go)
1. [compiler.Compile](../src/compiler/Compile.go)
1. [scanner.Scan](../src/scanner/Scan.go)
1. [core.Compile](../src/core/Compile.go)
1. [linker.Write](../src/linker/Write.go)

## FAQ

### Which platforms are supported?

|            | arm64 | x86-64 |
| ---------- | ----- | ------ |
| 🐧 Linux   | ✔️    | ✔️     |
| 🍏 Mac     | ✔️    | ✔️     |
| 🪟 Windows | ✔️    | ✔️     |

### How tiny is a Hello World?

|            | arm64      | x86-64    |
| ---------- | ---------- | --------- |
| 🐧 Linux   |  646 bytes | 646 bytes |
| 🍏 Mac     |   16.3 KiB |   4.2 KiB |
| 🪟 Windows |    1.7 KiB |   1.7 KiB |

### Are there any benchmarks?

Recursive Fibonacci benchmark (`n = 35`):

|                   | arm64                  | x86-64                 |
| ----------------- | ---------------------- | ---------------------- |
| C  (-O3, gcc 15)  | **41.4 ms** ±   1.4 ms | **26.2 ms** ±   4.1 ms |
| Q  (2025-08-20)   | **54.2 ms** ±   1.6 ms | **37.3 ms** ±   2.9 ms |
| Go (1.25, new GC) | **57.7 ms** ±   1.4 ms | **38.1 ms** ±   7.7 ms |
| C  (-O0, gcc 15)  | **66.4 ms** ±   1.5 ms | **52.3 ms** ±   5.2 ms |

While the current results lag behind optimized C, this is an expected stage of development. I am actively working to improve the compiler's code generation to a level that can rival optimized C, and I expect a significant performance uplift as this work progresses.

### What is the compiler based on?

The backend is built on a [Static Single Assignment (SSA)](https://en.wikipedia.org/wiki/Static_single-assignment_form) intermediate representation, the same approach used by mature compilers such as `gcc`, `go`, and `llvm`. SSA greatly simplifies the implementation of common optimization passes, allowing the compiler to produce relatively high-quality assembly code despite the project's early stage of development.

### Can I use it in scripts?

Yes. The compiler can build an entire script within a few microseconds.

```q
#!/usr/bin/env q
import io

main() {
	io.write("Hello\n")
}
```

You need to create a file with the contents above and add execution permissions via `chmod +x`. Now you can run the script without an explicit compiler build. The generated machine code runs directly from RAM if the OS supports it.

### Any security features?

**PIE**: All executables are built as position independent executables supporting a dynamic base address.

**W^X**: All memory pages are loaded with either execute or write permissions but never with both. Constant data is read-only.

|        | Read | Execute | Write |
| ------ | ---- | ------- | ----- |
| Code   | ✔️   | ✔️      | ❌    |
| Data   | ✔️   | ❌      | ❌    |

### Any editor extensions?

**Neovim**: Planned.

**VS Code**: Clone the [vscode-q](https://git.urbach.dev/extra/vscode-q) repository into your extensions folder (it enables syntax highlighting).

### Why is it written in Go and not language X?

Because of readability and great tools for concurrency.
The implementation will be replaced by a self-hosted compiler in the future.

### I can't contribute but can I donate to the project?

- [Kofi](https://ko-fi.com/akyoto)
- [GitHub](https://github.com/sponsors/akyoto)
- [Stripe](https://buy.stripe.com/4gw7vf5Jxflf83m7st)

### If I donate, what will my money be used for?

Testing infrastructure and support for existing and new architectures.

### How do I pronounce the name?

/ˈkjuː/ just like `Q` in the English alphabet.

## FAQ: Contributors

### Do you accept contributions?

Not at the moment. This project is currently part of a solo evaluation. Contributions will be accepted starting 2025-12-01.

### How do I run the tests?

```shell
# Run all tests:
go run gotest.tools/gotestsum@latest

# Generate coverage:
go test -coverpkg=./... -coverprofile=cover.out ./...

# View coverage:
go tool cover -func cover.out
go tool cover -html cover.out
```

### How do I run the benchmarks?

```shell
# Run compiler benchmarks:
go test ./tests -run '^$' -bench . -benchmem

# Run compiler benchmarks in single-threaded mode:
GOMAXPROCS=1 go test ./tests -run '^$' -bench . -benchmem

# Generate profiling data:
go test ./tests -run '^$' -bench . -benchmem -cpuprofile cpu.out -memprofile mem.out

# View profiling data:
go tool pprof --nodefraction=0.1 -http=:8080 ./cpu.out
go tool pprof --nodefraction=0.1 -http=:8080 ./mem.out
```

### Any debugging tools?

I recently added a preliminary `io.writeInt` to have some basic output for numeric values during compiler development.

You can also use the excellent [blinkenlights](https://justine.lol/blinkenlights/) from Justine Tunney to step through the x86-64 executables.

### Is there an IRC channel?

[#q](ircs://irc.urbach.dev:6697/#q) on [irc.urbach.dev](https://irc.urbach.dev).

## License

Please see the [license documentation](https://urbach.dev/license).

## Copyright

© 2025 Eduard Urbach