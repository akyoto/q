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
> `q` is under heavy development and not ready for production yet.
>
> Please read [this](https://lobste.rs/s/t7osqo/q_programming_language).
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

Run `hello` example:

```shell
q examples/hello
```

Build an executable:

```shell
q build examples/hello
```

Build for another architecture:

```shell
q build examples/hello --arch arm
q build examples/hello --arch x86
```

Build for another operating system:

```shell
q build examples/hello --os linux
q build examples/hello --os mac
q build examples/hello --os windows
```

## News

- 2025-08-12: Added support for Windows on arm64.
- 2025-08-11: Added support for Mac on arm64.

## Syntax

> [!NOTE]
> This is a draft.

[hello.q](../examples/hello/hello.q):

```
import io

main() {
	io.write("Hello\n")
}
```

[fibonacci.q](../examples/fibonacci/fibonacci.q)

```
fibonacci(n int) -> int {
	if n <= 1 {
		return n
	}

	return fibonacci(n - 1) + fibonacci(n - 2)
}
```

[gcd.q](../examples/gcd/gcd.q)

```
gcd(a int, b int) -> int {
	loop {
		switch {
			a == b { return a }
			a > b  { a -= b }
			_      { b -= a }
		}
	}
}
```

[fizzbuzz.q](../examples/fizzbuzz/fizzbuzz.q):

```
fizzbuzz(n int) {
	loop x := 1..n+1 {
		switch {
			x % 15 == 0 { io.write("FizzBuzz") }
			x % 5 == 0  { io.write("Buzz") }
			x % 3 == 0  { io.write("Fizz") }
			_           { io.writeInt(x) }
		}

		if x != n {
			io.write(" ")
		}
	}
}
```

The work is currently being focused on the correctness of the compiler and the proper code generation for all architectures and operating systems. The language syntax is highly volatile at this point but you can take a look at the [examples](../examples) or the [tests](../tests) to get a perspective on the current status. Documentation for all language features will follow once the core systems are stable.

## Source

This section is for contributors who want a high-level overview of the source code structure which uses a flat layout without nesting:

### Packages

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

### Typical flow

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
| 🐧 Linux   |  646 bytes | 582 bytes |
| 🍏 Mac     |   16.3 KiB |   4.2 KiB |
| 🪟 Windows |    1.7 KiB |   1.7 KiB |

This table often raises the question why Mac builds are so huge compared to the rest. The answer is in [these few lines](https://github.com/apple-oss-distributions/xnu/blob/e3723e1f17661b24996789d8afc084c0c3303b26/bsd/kern/mach_loader.c#L2021-L2027) of their kernel code. None of the other operating systems force you to page-align sections on disk. In practice, however, it's not as bad as it sounds because the padding is a zero-filled area that barely consumes any disk space in [sparse files](https://en.wikipedia.org/wiki/Sparse_file).

### How is the assembly code quality?

The backend uses an SSA based IR which is also used by well established compilers like `gcc`, `go` and `llvm`. SSA makes it trivial to apply lots of common optimization passes to it. As such, the quality of the generated assembly is fairly high despite the young age of the project.

### How do I use it for scripting?

The compiler is actually so fast that it's possible to compile an entire script within microseconds.

```q
#!/usr/bin/env q

import io

main() {
	io.write("Hello\n")
}
```

Create a file with the contents above and add permissions via `chmod +x`. Now you can execute it from anywhere. The generated machine code runs directly from RAM if the OS supports it.

### Which security features are supported?

#### PIE

All executables are built as position independent executables supporting a dynamic base address.

#### W^X

All memory pages are loaded with either execute or write permissions but never with both. Constant data is read-only.

|        | Read | Execute | Write |
| ------ | ---- | ------- | ----- |
| Code   | ✔️   | ✔️      | ❌    |
| Data   | ✔️   | ❌      | ❌    |

### Any editor extensions?

- VS Code (basic highlighting): Clone the [vscode-q](https://git.urbach.dev/extra/vscode-q) repository into your extensions folder
- Neovim support is planned.

### Why Go and not language X?

Because of readability, maintainability and great tools for concurrency.
The implementation will be replaced by a self-hosted compiler in the future.

### How do I pronounce the name?

/ˈkjuː/ just like `q` in the English alphabet.

### I can't contribute but can I donate to the project?

Yes, you can [donate](https://buy.stripe.com/4gw7vf5Jxflf83m7st) if you live in a country supported by Stripe.
The payment is classified as a "cash donation".

### If I donate, what will my money be used for?

Server infrastructure and support for existing and new architectures.

## FAQ: Contributors

### How do I run the test suite?

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

### Is there an IRC channel?

[#q](ircs://irc.urbach.dev:6697/#q) on [irc.urbach.dev](https://irc.urbach.dev).

## License

Please see the [license documentation](https://urbach.dev/license).

## Copyright

© 2025 Eduard Urbach