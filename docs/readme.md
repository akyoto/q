<div align="center">
	<img src="logo.svg" width="128" alt="q logo">
	<h1>The Q Programming Language</h1>
</div>

> [!NOTE]
> `q` is under heavy development and not ready for production yet.
>
> Feel free to [get in touch](https://urbach.dev/contact) if you are interested in helping out.

## Features

* High performance (`ssa` and `asm` optimizations)
* Tiny executables ("Hello World" is ~600 bytes)
* Fast compilation (<1 ms for simple programs)
* Unix scripting (pseudo JIT)
* No dependencies (no llvm, no libc)

## Installation

```shell
git clone https://git.urbach.dev/cli/q
cd q
go build
```

Add a symlink in `~/.local/bin` or `/usr/local/bin`:

```shell
ln -s $PWD/q ~/.local/bin/q
```

## Usage

Quick test:

```shell
q examples/hello
```

Build an executable:

```shell
q build examples/hello
```

Cross-compile for another OS:

```shell
q build examples/hello --os windows
```

## Tests

```shell
go run gotest.tools/gotestsum@latest
```

## Source

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

Those marked with a star need testing. Please get in touch if you have a machine with the marked architectures.

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

The compiler is actually so fast that it's possible to compile an entire script within microseconds. Create a new file...

```q
#!/usr/bin/env q

import io

main() {
	io.write("Hello\n")
}
```

...and add permissions via `chmod +x`. Now you can execute it from anywhere. The generated machine code runs directly from RAM if the OS supports it.

## License

Please see the [license documentation](https://urbach.dev/license).

## Copyright

© 2025 Eduard Urbach