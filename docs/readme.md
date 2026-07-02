<div align="center">
	<img src="logo.svg" width="90" alt="q logo">
	<h1>The Q Programming Language</h1>
	<p>
		<a href="#features">Features</a> |
		<a href="#motivation">Motivation</a> |
		<a href="#news">News</a> |
		<a href="#installation">Installation</a> |
		<a href="#usage">Usage</a> |
		<a href="#examples">Examples</a> |
		<a href="#reference">Reference</a> |
		<a href="#faq">FAQ</a>
	</p>
	<p>Q is a minimal, dependency-free programming language and compiler targeting x86-64 and arm64 with ultra-fast builds and tiny binaries.</p>
</div>

## Features

- ⚡ High performance (comparable to C and Go)
- 🚀 Fast compilation (5-10x faster than most compilers)
- 📦 Lightweight executables (1 KB for simple programs)
- 🔍 Static analysis (integrated linter catches common mistakes)
- 🛡️ Pointer safety (pointers cannot be nil)
- ♻️ Resource safety (use-after-free is a compile error)
- 🧠 Simple syntax (control flow is easily understood)
- 💬 Friendly errors (clear and concise compiler messages)
- 🌐 General purpose (apps, servers, games, kernels, etc.)
- 🧩 Multiple architectures (x86-64 and arm64)
- 🖥️ Multiple platforms (Linux, Mac and Windows)
- 📖 Readable source (less than 1% of LLVM's code size)
- 🧘 Zero dependencies (no external tools or libraries)

## Motivation

Q is a programming language under development that aims to fill the gap between C and Go while building upon the safety mechanisms that languages like Austral and Rust have demonstrated. Go's implementation details like the garbage collector make it difficult to be used in latency sensitive environments like kernel or game development and also complicate an efficient foreign function interface. However, the language itself is extremely well-designed when it comes to readability. Programs are written only once but they must be read many more times by developers around the globe and future generations trying to decipher the backbones of our software landscape. I want to combine the simplicity and readability of Go with a more low-level approach to memory safety so that we have a systems programming language that is both safe and efficient but also easy to understand for future readers.

Q is also a code generation framework that aims to produce raw machine code for multiple architectures, similar to LLVM. Since it is still a very young project compared to LLVM's 22 years of development, it will take time to reach a similar level of runtime performance for the generated executables. However, when looking at the compiler efficiency, the benchmarks show that many of the common compilers used in the industry are inefficient and that there is a lot of room for improvement. The Q compiler is currently the fastest optimizing compiler. While there are many optimization passes that still need to be implemented, I am confident that the performance impact of future passes can be reduced to a minimum. This project aims to raise the bar for compiler efficiency and demonstrate the possible improvements.

## Status

> [!WARNING]
> Q is [in early development](https://lobste.rs/s/t7osqo/q_programming_language) and not ready for production yet.
>
> The compiler currently passes a total of [2900 tests](#how-do-i-run-the-tests).
>
> Feel free to [contact me](https://urbach.dev/contact) if you are interested in contributing.

## News

- **2026-06-28**: Mutex synchronization.
- **2026-06-24**: Fast memory allocations.
- **2026-06-23**: Automatic deallocation.
- **2026-06-03**: Method calls.
- **2026-05-07**: Struct initialization.
- **2026-05-06**: Command line arguments.
- **2025-10-10**: Loop control flow.
- **2025-10-05**: Struct types in fields.
- **2025-09-30**: Static allocations.
- **2025-09-22**: Array allocations.
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

## Installation

### Build from source

If you have [Go](https://go.dev/) installed, you are 3 steps away from a working compiler:

```shell
git clone https://git.urbach.dev/cli/q
cd q
go build
```

### Install via symlink

The above steps produced the `q` binary. To access it from any location, add a symlink to a directory listed in your `PATH` environment variable:

```shell
ln -s $PWD/q ~/.local/bin/
```

## Usage

`q` expects a command as its first argument.

### run

If no command was specified, `q` defaults to the `run` command which executes the code in a source file.
Using it on directories is equivalent to listing the included `.q` files manually.

```shell
q examples/hello
```

### build

The `build` command produces an executable file inside the specified directory:

```shell
q build examples/hello
```

`q` implements its own assembler and linker. You can easily cross-compile for a different platform:

```shell
q build examples/hello -os windows
```

Leaving out the directory starts a build in the current directory:

```shell
q build
```

### ssa

Shows the [SSA](https://en.wikipedia.org/wiki/Static_single-assignment_form) form:

```shell
q ssa examples/hello
```

You can filter the output by function name with the `-func` option.

### asm

Shows the [assembly code](https://en.wikipedia.org/wiki/Assembly_language):

```shell
q asm examples/hello
```

You can filter the output by function name with the `-func` option.

## Examples

### hello

```q
import io

main() {
	io.write("Hello\n")
}
```

### echo

```q
echo() {
	buffer := new(byte, 4096)

	loop {
		n, _ := io.read(buffer)

		if n == 0 {
			return
		}

		io.write(buffer[..n])
	}
}
```

### fibonacci

```q
fibonacci(n int) -> int {
	if n <= 1 {
		return n
	}

	return fibonacci(n - 1) + fibonacci(n - 2)
}
```

### fizzbuzz

```q
fizzbuzz(x int) {
	switch {
		x % 15 == 0 { io.write("FizzBuzz") }
		x % 5 == 0  { io.write("Buzz") }
		x % 3 == 0  { io.write("Fizz") }
		_           { io.write(x) }
	}
}
```

See more in the [examples](../examples) directory.

## Reference

The following is a cheat sheet documenting the syntax.

| I need to...                         |                              | Stability       |
| ------------------------------------ | ---------------------------- | --------------- |
| Define a new variable                | `x := 1`                     | ✔️ Stable       |
| Reassign an existing variable        | `x = 2`                      | ✔️ Stable       |
| Define a function                    | `main() {}`                  | ✔️ Stable       |
| Define a struct                      | `Point {}`                   | ✔️ Stable       |
| Define input and output types        | `f(a int) -> (b int) {}`     | ✔️ Stable       |
| Define same function for other types | `f(_ string) {} f(_ int) {}` | 🚧 Experimental |
| Instantiate a struct                 | `Point{x: 1, y: 2}`          | ✔️ Stable       |
| Allocate a type                      | `new(int)`                   | 🚧 Experimental |
| Allocate an array                    | `new(int, 10)`               | 🚧 Experimental |
| Allocate a struct                    | `new(Point)`                 | 🚧 Experimental |
| Allocate and initialize a struct     | `new(Point){x: 1, y: 2}`     | 🚧 Experimental |
| Delete an object                     | `delete(p)`                  | ✔️ Stable       |
| Define a struct method               | `f(p *Point) {}`             | 🚧 Experimental |
| Call struct methods                  | `p.f()`                      | ✔️ Stable       |
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
| Jump to the next iteration           | `loop.next()`                | 🚧 Experimental |
| Jump to the end of the loop          | `loop.stop()`                | 🚧 Experimental |
| Branch                               | `if {} else {}`              | ✔️ Stable       |
| Branch multiple times                | `switch { cond {} _ {} }`    | ✔️ Stable       |
| Define a constant                    | `const { x = 42 }`           | ✔️ Stable       |
| Declare an external function         | `extern { g { f() } }`       | ✔️ Stable       |
| Output a string                      | `io.write("Hello\n")`        | ✔️ Stable       |
| Output an integer                    | `io.write(42)`               | ✔️ Stable       |
| Cast a type                          | `x as byte`                  | 🚧 Experimental |
| Mark a type as a resource            | `!`                          | 🚧 Experimental |
| Mark a parameter as unused           | `_`                          | ✔️ Stable       |

### Tokens

Source files are preprocessed by the tokenizer which groups the individual bytes into several token types:

- [Identifier](#identifiers)
- [Number](#numbers)
- [Rune](#runes)
- [String](#strings)
- [Comment](#comments)
- [Operators](#operators)
- [Keywords](#keywords)
- [Builtins](#builtins)

### Identifiers

An identifier like `x` is a non-empty sequence of letters, digits, and underscores (`_`).
The first character of an identifier must not be a digit. Identifiers are case-sensitive.

### Numbers

A number like `42` is a non-empty sequence of digits. It may start with a `-` to indicate negative values. Numbers are decimal by default but the base can be overriden with a `0x` prefix for hexadecimal, `0o` for octal and `0b` for binary. The uppercase letters from `A` to `F` are used to represent digits from 10 to 15 in hexadecimal.

### Runes

A rune literal like `'日'` or `'本'` is an integer representing a Unicode code point. It must be enclosed by `'`. It is equivalent to an integer from the perspective of the compiler. The value of the integer is derived from the Unicode representation of the content.

```q
assert 'A' == 0x41
assert 'a' == 0x61
assert '世' == 0x4E16
assert '界' == 0x754C
assert '😀' == 0x1F600
```

### Strings

A string literal like `"Hello"` is a sequence of bytes enclosed by `"`. Strings are immutable, though the compiler does not enforce this rule in its present state. The following escape sequences starting with `\` can be used in rune and string literals to embed special characters:

```
assert '\0' == 0
assert '\t' == 9
assert '\n' == 10
assert '\r' == 13
assert '\"' == 34
assert '\'' == 39
assert '\\' == 92
```

### Comments

A line comment like `// This is a comment` starts with `//` and stops at the end of the line. Comments are ignored by the compiler and can be added to the code for documentation purposes. Multiline comments are not supported.

### Operators

Operators like `+` represent binary or unary operations.

Operator precedence defines the order of operations. An operation with a higher precedence is performed before operations with lower precedence. Precedence levels introduce additional rules that programmers must learn and can lead to hidden mistakes.
To minimize this complexity, Q is limiting the operators to only 8 precedence levels:

| Precedence | Operators                                                         | Description             |
| ---------: | ----------------------------------------------------------------- | ----------------------- |
|          8 | `.` `()` `[]` `{}`                                                | Postfix                 |
|          7 | `!` `-`                                                           | Unary                   |
|          6 | `*` `/` `%`                                                       | Multiplicative          |
|          5 | `+` `-` `&` `\|` `^` `<<` `>>` `as`                               | Additive, bitwise, cast |
|          4 | `==` `!=` `<` `>` `<=` `>=`                                       | Comparison              |
|          3 | `&&` `\|\|`                                                       | Logical                 |
|          2 | `..` `,`                                                          | Range, separator        |
|          1 | `:=` `=` `+=` `-=` `*=` `/=` `%=` `&=` `\|=` `^=` `<<=` `>>=` `:` | Assignment              |

### Keywords

| Keyword    | Description                                                | Stability       |
| ---------- | ---------------------------------------------------------- | --------------- |
| `assert`   | Tests conditions at runtime                                | ✔️ Stable       |
| `const`    | Defines constant expressions                               | ✔️ Stable       |
| `else`     | Failure branch for if statements                           | ✔️ Stable       |
| `extern`   | Foreign function definitions                               | ✔️ Stable       |
| `global`   | Global variables (discouraged but required in stdlib)      | ✔️ Stable       |
| `go`       | Asynchronous function calls                                | 🚧 Experimental |
| `if`       | Branches based on a condition                              | ✔️ Stable       |
| `import`   | Allows access to other packages                            | ✔️ Stable       |
| `loop`     | Repeatable code                                            | ✔️ Stable       |
| `return`   | Ends the function and returns values to the caller         | ✔️ Stable       |
| `switch`   | Multiple branches executing the first true condition block | ✔️ Stable       |

### Builtins

| Function   | Description                                                | Stability       |
| ---------- | ---------------------------------------------------------- | --------------- |
| `cas`      | Atomic compare and swap                                    | 🚧 Experimental |
| `delete`   | Frees memory                                               | ✔️ Stable       |
| `new`      | Allocates memory                                           | 🚧 Experimental |
| `syscall`  | Calls a kernel function                                    | ✔️ Stable       |

### Packages

A package is defined by a directory.
All files in that directory belong to the same package and share access to its identifiers.
Subdirectories form separate packages.

Within each package, a function named `init` is executed automatically at program startup,
while a function named `exit` runs before the program terminates.
This feature is intended primarily for the standard library and is generally discouraged in application code.

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

For memory allocations of slices and pointers `delete` is called automatically on all exit points of the identifier's scope.
Non-pointer types like `!int` currently do not support automatic lifecycle management and require an explicit free call, e.g. an `fs.close` for `!int` file handles.

## Errors

> [!NOTE]
> Algebraic data types for error handling will be considered at a later point but as of now there are no final decisions on the matter.

Any function can currently define an `error` type return value at the end:

```
a, b, err := canFail()
```

An error value protects all the return values to the left of it.
The protected values `a` and `b` can not be accessed without checking `err` first.
Additionally, error variables like `err` are invalidated after the branch that checked them.

```
a, b, err := canFail()

// ❌ a and b are inaccessible
// ✅ err is accessible

if err != 0 {
	return
}

// ✅ a and b are accessible
// ❌ err is no longer accessible
```

The `error` type is currently defined to be an integer, though this is expected to change in a future version.

## Security

Recent incidents such as the `xz` backdoor and attacks on the `npm` ecosystem have shown that supply chain attacks remain one of the software industry's biggest security challenges.

Q helps mitigate these risks by enforcing the principle of least privilege. Every module must explicitly declare which sensitive resources it requires, such as network or file system access. Any permission changes become part of the review process during updates, making unexpected behavior much more visible. If `leftpad` suddenly requests access to `net`, that should immediately raise suspicion.

While this cannot eliminate supply chain attacks entirely, it significantly reduces the chances of your system being compromised.

Q also hardens executables at the binary level:

- All executables are built as position-independent executables (PIE) with dynamic base addresses so that an attacker can't use precalculated addresses.
- The call stack where return addresses are located is isolated from the regular memory stack, eliminating an entire class of control-flow attacks.
- The W^X (write xor execute) policy is enforced for all memory pages: memory can be writable or executable, but never both.

## Syntax

Q encourages code editors to implement multiple syntaxes for editing.

A view of the code can be substantially different from the underlying model that is saved to disk.
It's important to conceptually realize that one is just a temporary view for editing and the other is a form of persistent data storage.

It is absolutely possible that an editor could offer editing in a Python-like whitespace-significant view.
It is also possible to offer visual editing with a node-based system similar to Scratch or Unreal Engine blueprints.
In all cases the code that is saved to disk would still use the standard text-based format.

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
- [optimizer](../src/optimizer) - Code optimization
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

An online view of analytic tools can be found here:

- [Dependency Graph](https://deps-q.urbach.dev/) via gopkgview
- [Flame Graph](https://prof-q.urbach.dev/ui/flamegraph) via pprof

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
| 🍏 Mac     | 32.3 KiB | 8.2 KiB |
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

### Any editor extensions?

- ~~**Neovim**~~: Planned.
- **VS Code**: Clone the [vscode-q](https://git.urbach.dev/extra/vscode-q) repository into your extensions folder (it enables syntax highlighting).
- ~~**Zed**~~: Planned.

### Why is it written in Go and not language X?

Because of readability and great tools for concurrency.
The implementation will be replaced by a self-hosted compiler in the future.

### I can't contribute but can I donate to the project?

- [Bmac](https://buymeacoffee.com/akyoto)
- [GitHub](https://github.com/sponsors/akyoto)
- [Kofi](https://ko-fi.com/akyoto)

### If I donate, what will my money be used for?

Continuous development, testing infrastructure and support for new platforms and architectures.

### How do I pronounce the name?

/ˈkjuː/ just like `Q` in the English alphabet.

## FAQ: Contributors

### How do I run the tests?

```shell
# Run all tests
go run gotest.tools/gotestsum@latest

# Generate coverage
go test -coverpkg=./... -coverprofile=cover.out ./...

# View coverage
go tool cover -func cover.out
go tool cover -html cover.out
```

### How do I run the benchmarks?

```shell
# Run compiler benchmarks
go test ./tests -run '^$' -bench . -benchmem

# Run compiler benchmarks in single-threaded mode
GOMAXPROCS=1 go test ./tests -run '^$' -bench . -benchmem

# Generate profiling data
go test ./tests -run '^$' -bench . -benchmem -cpuprofile cpu.out -memprofile mem.out

# View profiling data
go tool pprof -http=:8080 ./cpu.out
go tool pprof -http=:8080 ./mem.out
```

### How do I run a single file in `tests`?

To run a single test file, linter errors must be disabled using the `-no-lint` flag:

```shell
q tests/add.q -no-lint
```

This is needed because tests often assert "obvious" facts that the linter would not allow in normal programs.

### How do I analyze a problem with the compiler?

Replace `q build` with `q ssa` or `q asm` to see the intermediate stages which reveal how the compiler understands your program code.
Use `-func` to filter out specific functions.

If that doesn't reveal any bugs, you can also use the excellent [blinkenlights](https://justine.lol/blinkenlights/) from Justine Tunney to step through x86-64 executables one instruction at a time.

### Is there a community?

* **IRC**: [#q](ircs://irc.urbach.dev:6697/#q) on [irc.urbach.dev](https://irc.urbach.dev) is the main hub for collaboration.
* **Discord**: [Q community](https://discord.gg/4q3DJFsTvB) is a more laid-back alternative that is popular among gamers.
* ~~**Forum**:~~ Web forums are currently not available.
* ~~**E-Mail**:~~ Mailing lists are currently not available.

## Thanks

In alphabetical order:

- [Andrew Binstock](https://github.com/platypusguy) | offer to help testing and documenting
- [Anto "xplshn"](https://github.com/xplshn) | feedback on public compiler interfaces
- [Bjorn De Meyer](https://github.com/bjorndm) | feedback on language features
- [James Mills](https://github.com/prologic) | first one to contact me about the project
- [Laurent Demailly](https://github.com/ldemailly) | indispensable help with Mac debugging
- [Max van IJsselmuiden](https://github.com/maxvij) | feedback and Mac debugging
- [Mustafa F. Bulut](https://github.com/mfbulut) | first one to buy me a coffee
- [Nikita Proskourine](https://github.com/Deobfuscator) | first monthly supporter on GitHub
- [Tibor Halter](https://github.com/zupa-hu) | detailed feedback and bug reporting

## License

Please see the [license documentation](https://urbach.dev/license).

## Copyright

© 2025 Eduard Urbach