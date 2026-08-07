# Source

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
- [data](../src/data) - Data container that can reuse existing data
- [dll](../src/dll) - DLL support for Windows systems
- [elf](../src/elf) - ELF format for Linux executables
- [errors](../src/errors) - Error handling that reports lines and columns
- [exe](../src/exe) - Generic executable format to calculate section offsets
- [expression](../src/expression) - Expression parser generating trees
- [fs](../src/fs) - File system access
- [global](../src/global) - Global variables like the working directory
- [linker](../src/linker) - Frontend for generating executable files
- [linter](../src/linter) - Linter that catches common mistakes
- [macho](../src/macho) - Mach-O format for Mac executables
- [memfile](../src/memfile) - Memory backed file descriptors
- [optimizer](../src/optimizer) - Code optimization
- [pe](../src/pe) - PE format for Windows executables
- [resolver](../src/resolver) - Type token resolver
- [scanner](../src/scanner) - Scanner that parses top-level instructions
- [set](../src/set) - Generic set implementation
- [ssa](../src/ssa) - Static single assignment types
- [token](../src/token) - Tokenizer
- [types](../src/types) - Type system
- [verbose](../src/verbose) - Verbose output
- [x86](../src/x86) - x86-64 architecture

## Flow

The typical flow for a build command is the following:

1. [main](../main.go)
2. [cli.Exec](../src/cli/Exec.go)
3. [compiler.Compile](../src/compiler/Compile.go)
4. [scanner.Scan](../src/scanner/Scan.go)
5. [core.Compile](../src/core/Compile.go)
6. [linker.Write](../src/linker/Write.go)

## Tools

- [Dependency Graph](https://deps-q.urbach.dev/) (gopkgview)
- [Flame Graph](https://prof-q.urbach.dev/ui/flamegraph) (pprof)
