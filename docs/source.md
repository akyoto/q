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

## Static Single Assignment

The SSA IR follows a simple rule: every value is assigned exactly once.

### Basic Blocks

Every function has a list of basic blocks. Basic blocks store values in the order they appear in the original code.

### Instructions

Instructions are no different from values. The calculation `1 + 2` is represented as 3 values using 2 int constants and 1 binary operation:

```
t0 = 1
t1 = 2
t2 = t0 + t1
```

All of these are considered values even in cases where the type of the value is `void` such as in procedural function calls with side effects.

### Pointers

There are no IDs or indices, therefore values don't know their position in a basic block.

Values reference other values by including a pointer to them.

## Executable formats

### Linux (ELF)

#### Basic structure

1. ELF header [0x00 : 0x40]
2. Program headers
3. String table
4. Section headers
5. Padding
6. Executable code
7. Padding
8. Read-only data

#### Entry point

The executables are compiled as position-independent executables (PIE).
Therefore the entry point is defined as a file offset instead of a static virtual address.

#### Padding

Permissions like read, write and execute can only be applied to an entire page in memory.
To ensure that execution permissions are properly applied, the code section and the data section are aligned on page boundaries.

#### Links

- https://git.kernel.org/pub/scm/linux/kernel/git/torvalds/linux.git/tree/fs/binfmt_elf.c
- https://git.kernel.org/pub/scm/linux/kernel/git/torvalds/linux.git/tree/arch/x86/include/asm/elf.h
- https://git.kernel.org/pub/scm/linux/kernel/git/torvalds/linux.git/tree/include/uapi/linux/elf.h
- https://lwn.net/Articles/631631/
- https://en.wikipedia.org/wiki/Executable_and_Linkable_Format
- https://www.muppetlabs.com/~breadbox/software/tiny/teensy.html
- https://nathanotterness.com/2021/10/tiny_elf_modernized.html

### Mac (Mach-O)

#### Notes

- The start of the file must be loaded in some segment.
- The start of the file must be marked as readable + executable.
- Load command size must be divisible by 8.
- Segments must be page-aligned in the file.

#### Links

- https://github.com/apple-oss-distributions/xnu/blob/main/bsd/kern/mach_loader.c
- https://github.com/apple-oss-distributions/xnu/blob/main/bsd/kern/mach_loader.c#L2021-L2027
- https://github.com/apple-oss-distributions/xnu/blob/main/EXTERNAL_HEADERS/mach-o/loader.h
- https://en.wikipedia.org/wiki/Mach-O
- https://github.com/aidansteele/osx-abi-macho-file-format-reference
- https://stackoverflow.com/questions/39863112/what-is-required-for-a-mach-o-executable-to-load

### Windows (PE)

#### Notes

Unlike Linux, Windows does not ignore zero-length sections at the end of a file and will
fail loading them because they don't exist within the file. Adding a single byte to the
section can fix this problem, but it's easier to just remove the section header entirely.
The solution used here is to guarantee that the data section is never empty by always
importing a few core functions from "kernel32.dll".

#### DLL function pointers

The section where the DLL function pointers are stored does not need to be marked as writable.
The Windows executable loader resolves the pointers before they are loaded into memory.

The stack must be 16 byte aligned before a DLL function is called.

#### Links

- https://learn.microsoft.com/en-us/windows/win32/debug/pe-format
- https://learn.microsoft.com/en-us/previous-versions/ms809762(v=msdn.10)
- https://learn.microsoft.com/en-us/archive/msdn-magazine/2002/february/inside-windows-win32-portable-executable-file-format-in-detail
- https://learn.microsoft.com/en-us/archive/msdn-magazine/2002/march/inside-windows-an-in-depth-look-into-the-win32-portable-executable-file-format-part-2
- https://blog.kowalczyk.info/articles/pefileformat.html
- https://keyj.emphy.de/win32-pe/
- https://corkamiwiki.github.io/PE
- https://github.com/ayaka14732/TinyPE-on-Win10