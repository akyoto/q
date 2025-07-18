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
* Tiny executables ("Hello World" is ~500 bytes)
* Fast compilation (much faster than other compilers)
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

### Unix scripts

The compiler is actually so fast that it's possible to use `q` for scripting. Create a new file...

```q
#!/usr/bin/env q

import io

main() {
	io.write("Hello\n")
}
```

...and add exec permissions via `chmod +x`. Now you can execute it from anywhere. The generated machine code is run directly from RAM if the OS supports it.

## Tests

```shell
go run gotest.tools/gotestsum@latest
```

## Source overview

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

## Platforms

|         | arm64  | x86-64 |
| ------- | ------ | ------ |
| Linux   | ✔️     | ✔️     |
| Mac     | ✔️*    | ✔️     |
| Windows | ✔️*    | ✔️     |

Those marked with a star are supported in theory but there are no developer machines to test them.

## Security

### PIE

All executables are built as Position Independent Executables (PIE) supporting a dynamic base address.

### Memory pages

Code and data are separated into different memory pages and loaded with different access permissions.

|        | Read | Execute | Write |
| ------ | ---- | ------- | ----- |
| Code   | ✔️   | ✔️      | ❌    |
| Data   | ✔️   | ❌      | ❌    |

## License

Please see the [license documentation](https://urbach.dev/license).

## Copyright

© 2025 Eduard Urbach