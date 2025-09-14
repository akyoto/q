<div align="center">
	<img src="logo.svg" width="90" alt="q logo">
	<h1>The Q Programming Language</h1>
	<p>Q is a minimal, dependency-free programming language and compiler targeting x86-64 and arm64 with ultra-fast builds and tiny binaries.</p>
</div>

## Features

- High performance (comparable to C and Go)
- Fast compilation (5-10x faster than most)
- Tiny executables ("Hello World" is 0.6 KiB)
- Static analysis (no need for external linters)
- Pointer safety (pointers cannot be nil)
- Resource safety (use-after-free is a compile error)
- Multiple platforms (Linux, Mac and Windows)
- Zero dependencies (no llvm, no libc)

## Installation

> [!WARNING]
> Q is [still in development](https://git.urbach.dev/cli/q/issues/1) and not ready for production yet.
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

- **2025-09-09**: Type casts.
- **2025-09-08**: Function pointers.
- **2025-09-07**: Pointer safety.
- **2025-09-03**: Error handling.
- **2025-08-31**: Constant folding.
- **2025-08-25**: Resource safety.
- **2025-08-23**: Function overloading.
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

A few selected examples:

- [hello](../examples/hello/hello.q)
- [echo](../examples/echo/echo.q)
- [fibonacci](../examples/fibonacci/fibonacci.q)
- [fizzbuzz](../examples/fizzbuzz/fizzbuzz.q)

Advanced examples using unstable APIs:

- [raylib](../examples/raylib/raylib.q)
- [server](../examples/server/server.q)

## Reference

The following is a cheat sheet documenting the syntax.

| I need to...                         |                              | API stability   |
| ------------------------------------ | ---------------------------- | --------------- |
| Define a new variable                | `x := 1`                     | ✔️ Stable       |
| Reassign an existing variable        | `x = 2`                      | ✔️ Stable       |
| Define a function                    | `main() {}`                  | ✔️ Stable       |
| Define a struct                      | `Point {}`                   | ✔️ Stable       |
| Define input and output types        | `f(a int) -> (b int) {}`     | ✔️ Stable       |
| Define same function for other types | `f(_ string) {} f(_ int) {}` | 🚧 Experimental |
| Instantiate a struct                 | `Point{x: 1, y: 2}`          | ✔️ Stable       |
| Instantiate a type on the heap       | `new(Point)`                 | 🚧 Experimental |
| Delete a type from the heap          | `delete(p)`                  | 🚧 Experimental |
| Access struct fields                 | `p.x`                        | ✔️ Stable       |
| Dereference a pointer                | `[ptr]`                      | ✔️ Stable       |
| Index a pointer                      | `ptr[0]`                     | ✔️ Stable       |
| Slice a string                       | `"Hello"[1..3]`              | ✔️ Stable       |
| Slice a string from index            | `"Hello"[1..]`               | ✔️ Stable       |
| Slice a string until index           | `"Hello"[..3]`               | ✔️ Stable       |
| Return multiple values               | `return 1, 2`                | ✔️ Stable       |
| Loop                                 | `loop {}`                    | ✔️ Stable       |
| Loop 10 times                        | `loop 0..10 {}`              | ✔️ Stable       |
| Loop 10 times with a variable        | `loop i := 0..10 {}`         | ✔️ Stable       |
| Branch                               | `if {} else {}`              | ✔️ Stable       |
| Branch multiple times                | `switch { cond {} _ {} }`    | ✔️ Stable       |
| Define a constant                    | `const { x = 42 }`           | ✔️ Stable       |
| Declare an external function         | `extern { g { f() } }`       | ✔️ Stable       |
| Allocate memory                      | `mem.alloc(4096)`            | ✔️ Stable       |
| Free memory                          | `mem.free(buffer)`           | 🚧 Experimental |
| Output a string                      | `io.write("Hello\n")`        | ✔️ Stable       |
| Output an integer                    | `io.write(42)`               | ✔️ Stable       |
| Cast a type                          | `x as byte`                  | 🚧 Experimental |
| Mark a type as a resource            | `!`                          | 🚧 Experimental |
| Mark a parameter as unused           | `_`                          | ✔️ Stable       |

## Resources

> [!WARNING]
> This feature is very new and still undergoing refinement.
> For more information, refer to linear types or borrowing in other languages.

Resources are shared objects such as files, memory or network sockets. The use of resource types prevents the following problems:

- **Resource leaks** (forgetting to free a resource)
- **Use-after-free** (using a resource after it was freed)
- **Double-free** (freeing a resource twice)

Any type, even integers, can be turned into a resource by prefixing the type with `!`. For example, consider these minimal functions:

```
alloc() -> !int { return 1 }
use(_ int) {}
free(_ !int) {}
```

With this, forgetting to call `free` becomes impossible:

```
x := alloc()
use(x)
```

```
x := alloc()
     ┬
     ╰─ Resource of type '!int' not consumed
```

Attempting a use-after-free is also rejected:

```
x := alloc()
free(x)
use(x)
```

```
use(x)
    ┬
    ╰─ Unknown identifier 'x'
```

Likewise, a double-free is disallowed:

```
x := alloc()
free(x)
free(x)
```

```
free(x)
free(x)
     ┬
     ╰─ Unknown identifier 'x'
```

The compiler only accepts the correct usage order:

```
x := alloc()
use(x)
free(x)
```

The `!` prefix marks a type to be consumed exactly once. It has no runtime overhead. When a `!int` is passed to another `!int`, the original variable is invalidated in subsequent code. As an exception, converting `!int` to `int` bypasses this rule, allowing multiple uses.

The standard library currently makes use of this feature in two packages:

- `fs.open` must be followed by `fs.close`
- `mem.alloc` must be followed by `mem.free`

## Errors

Any function can define an `error` type return value at the end:

```
a, b, err := canFail()
```

An error value protects all the return values to the left of it.
The protected values `a` and `b` can not be accessed without checking `err` first.
Additionally, error variables like `err` are invalidated after the branch that checked them.

```
a, b, err := canFail()

// a and b are inaccessible

if err != 0 {
	return
}

// a and b are accessible
// err is no longer defined
```

The `error` type is currently defined to be an integer. This will most likely change in a future version.

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
- [fold](../src/fold) - Constant folding
- [fs](../src/fs) - File system access
- [global](../src/global) - Global variables like the working directory
- [linker](../src/linker) - Frontend for generating executable files
- [macho](../src/macho) - Mach-O format for Mac executables
- [memfile](../src/memfile) - Memory backed file descriptors
- [pe](../src/pe) - PE format for Windows executables
- [scanner](../src/scanner) - Scanner that parses top-level instructions
- [set](../src/set) - Generic set implementation
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

There is also an interactive [dependency graph](https://deps-q.urbach.dev/) and a [flame graph](https://prof-q.urbach.dev/ui/flamegraph) (via gopkgview and pprof).

## FAQ

### Which platforms are supported?

|            | arm64 | x86-64 |
| ---------- | ----- | ------ |
| 🐧 Linux   | ✔️    | ✔️     |
| 🍏 Mac     | ✔️    | ✔️     |
| 🪟 Windows | ✔️    | ✔️     |

### How tiny is a Hello World?

|            | arm64    | x86-64  |
| ---------- | -------- | ------- |
| 🐧 Linux   |  0.6 KiB | 0.6 KiB |
| 🍏 Mac     | 16.3 KiB | 4.2 KiB |
| 🪟 Windows |  1.7 KiB | 1.7 KiB |

### Are there any runtime benchmarks?

Recursive Fibonacci benchmark (`n = 35`):

|                   | arm64                | x86-64               |
| ----------------- | -------------------- | -------------------- |
| C (-O3, gcc 15)   | **41.4 ms** ± 1.4 ms | **24.5 ms** ± 3.2 ms |
| Q (2025-08-20)    | **54.2 ms** ± 1.6 ms | **34.8 ms** ± 2.3 ms |
| Go (1.25, new GC) | **57.7 ms** ± 1.4 ms | **37.9 ms** ± 6.9 ms |
| C (-O0, gcc 15)   | **66.4 ms** ± 1.5 ms | **47.8 ms** ± 4.4 ms |

While the current results lag behind optimized C, this is an expected stage of development. I am actively working to improve the compiler's code generation to a level that can rival optimized C, and I expect a significant performance boost as this work progresses.

### Are there any compiler benchmarks?

The table below shows latency numbers on a 2015 Macbook:

|                 | x86-64                  |
| --------------- | ----------------------- |
| q               |   **78.6 ms** ±  2.3 ms |
| go @1.25        |  **364.5 ms** ±  3.3 ms |
| clang @17.0.0   |  **395.9 ms** ±  3.3 ms |
| rustc @1.89.0   |  **639.9 ms** ±  3.1 ms |
| v @0.4.11       | **1117.0 ms** ±  3.0 ms |
| zig @0.15.1     | **1315.0 ms** ± 12.0 ms |
| odin @accdd7c2a | **1748.0 ms** ±  8.0 ms |

Latency measures the time it takes a compiler to create an executable file with a nearly empty main function. It should not be confused with throughput.

Advanced benchmarks for throughput have not been conducted yet, but the following table shows timings in an extremely simplified test parsing 1000 Fibonacci functions named `fib0` to `fib999`:

|                 | x86-64                  |
| --------------- | ----------------------- |
| q               |   **89.5 ms** ±  2.4 ms |
| go @1.25        |  **372.2 ms** ±  5.3 ms |
| clang @17.0.0   |  **550.8 ms** ±  3.8 ms |
| rustc @1.89.0   | **1101.0 ms** ±  4.0 ms |
| v @0.4.11       | **1256.0 ms** ±  4.0 ms |
| zig @0.15.1     | **1407.0 ms** ± 12.0 ms |
| odin @accdd7c2a | **1770.0 ms** ±  7.0 ms |

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

### Does it have a fast memory allocator?

No, the current implementation is only temporary and it needs to be replaced with a faster one once the required language features have been implemented.

### Any security features?

**PIE**: All executables are built as position independent executables supporting a dynamic base address.

**W^X**: All memory pages are loaded with either execute or write permissions but never with both. Constant data is read-only.

|      | Read | Execute | Write |
| ---- | ---- | ------- | ----- |
| Code | ✔️   | ✔️      | ❌    |
| Data | ✔️   | ❌      | ❌    |

### Any editor extensions?

**Neovim**: Planned.

**VS Code**: Clone the [vscode-q](https://git.urbach.dev/extra/vscode-q) repository into your extensions folder (it enables syntax highlighting).

### Why is it written in Go and not language X?

Because of readability and great tools for concurrency.
The implementation will be replaced by a self-hosted compiler in the future.

### I can't contribute but can I donate to the project?

- [Kofi](https://ko-fi.com/akyoto)
- [GitHub](https://github.com/sponsors/akyoto)
- [Stripe](https://donate.stripe.com/4gw7vf5Jxflf83m7st)

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

The compiler has verbose output showing how it understands the program code using `--ssa` and `--asm` flags.
If that doesn't reveal any bugs, you can also use the excellent [blinkenlights](https://justine.lol/blinkenlights/) from Justine Tunney to step through the x86-64 executables one instruction at a time.

### Is there an IRC channel?

[#q](ircs://irc.urbach.dev:6697/#q) on [irc.urbach.dev](https://irc.urbach.dev).

## Thanks

In alphabetical order:

- [Anto "xplshn"](https://github.com/xplshn) | feedback on public compiler interfaces
- [Bjorn De Meyer](https://github.com/bjorndm) | feedback on PL design
- [Furkan](https://github.com/mfbulut) | first one to buy me a coffee
- [James Mills](https://github.com/prologic) | first one to contact me about the project
- [Laurent Demailly](https://github.com/ldemailly) | indispensable help with Mac debugging
- [Max van IJsselmuiden](https://github.com/maxvij) | feedback and Mac debugging
- [Nikita Proskourine](https://github.com/Deobfuscator) | first monthly supporter on GitHub
- [Tibor Halter](https://github.com/zupa-hu) | detailed feedback and bug reporting
- my wife :) | providing syntax feedback as a non-programmer

## License

Please see the [license documentation](https://urbach.dev/license).

## Copyright

© 2025 Eduard Urbach