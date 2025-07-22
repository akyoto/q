<div align="center">
	<img src="logo.svg" width="128" alt="q logo">
	<h1>The Q Programming Language</h1>
	<p>Q is a minimal, dependency-free programming language and compiler targeting x86-64 and arm64 with ultra-fast builds and tiny binaries.</p>
</div>

## Features

* High performance (`ssa` and `asm` optimizations)
* Fast compilation (<1 ms for simple programs)
* Tiny executables ("Hello World" is ~600 bytes)
* Many platforms (Linux, Mac and Windows)
* Zero dependencies (no llvm, no libc)

## Installation

> [!NOTE]
> `q` is under heavy development and not ready for production yet.
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

Show compiler output:

```shell
q build examples/hello --verbose
```

Cross-compile for another OS:

```shell
q build examples/hello --os linux
q build examples/hello --os mac
q build examples/hello --os windows
```

## Tests

Run all tests:

```shell
go run gotest.tools/gotestsum@latest
```

Generate coverage:

```shell
go test -coverpkg=./... -coverprofile=cover.out ./...
```

View coverage:

```shell
go tool cover -func cover.out
go tool cover -html cover.out
```

Run compiler benchmarks:

```shell
go test ./tests -run='^$' -bench=. -benchmem
```

Generate profiling data:

```shell
go test ./tests -run='^$' -bench="Examples/" -benchmem -cpuprofile cpu.out -memprofile mem.out
```

View profiling data:

```shell
go tool pprof --nodefraction=0.1 -http=:8080 ./cpu.out
go tool pprof --nodefraction=0.1 -http=:8080 ./mem.out
```

## Source overview

This section is for contributors who want a high-level overview of the source code structure.

### Packages

- [arm](../src/arm) - arm64 architecture
- [asm](../src/asm) - Generic assembler
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
- [x86](../src/x86) - x86-64 architecture

### Typical flow

1. [main](../main.go)
1. [cli.Exec](../src/cli/Exec.go)
1. [compiler.Compile](../src/compiler/Compile.go)
1. [scanner.Scan](../src/scanner/Scan.go)
1. [core.Compile](../src/core/Compile.go)
1. [linker.Write](../src/linker/Write.go)

## FAQ

### How tiny is a Hello World?

|         | x86-64 |
| ------- | ------ |
| Linux   | 582 bytes  |
| Mac     | 8 KiB     |
| Windows | 1.7 KiB     |

### How is the assembly code quality?

The backend uses an SSA based IR which is also used by well established compilers like `gcc`, `go` and `llvm`. SSA makes it trivial to apply lots of common optimization passes to it. As such, the quality of the generated assembly is pretty high despite the young age of the project.

### Which platforms are supported?

|         | arm64  | x86-64 |
| ------- | ------ | ------ |
| Linux   | ✔️     | ✔️     |
| Mac     | ✔️*    | ✔️     |
| Windows | ✔️*    | ✔️     |

Those marked with a star need testing. Please contact me if you have a machine with the marked architectures.

### Which security features are supported?

#### PIE

All executables are built as position independent executables supporting a dynamic base address.

#### W^X

All memory pages are loaded with either execute or write permissions but never with both. Constant data is read-only.

|        | Read | Execute | Write |
| ------ | ---- | ------- | ----- |
| Code   | ✔️   | ✔️      | ❌    |
| Data   | ✔️   | ❌      | ❌    |

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

### How do I pronounce the name?

/ˈkjuː/ just like `Q` in the English alphabet.

## License

Please see the [license documentation](https://urbach.dev/license).

## Copyright

© 2025 Eduard Urbach