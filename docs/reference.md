# Reference

## Cheatsheet

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

## Tokens

Source files are preprocessed by the tokenizer which groups the individual bytes into several token types:

- [Identifier](#identifiers)
- [Number](#numbers)
- [Rune](#runes)
- [String](#strings)
- [Comment](#comments)
- [Operators](#operators)
- [Keywords](#keywords)
- [Builtins](#builtins)

## Identifiers

An identifier like `x` is a non-empty sequence of letters, digits, and underscores (`_`).
The first character of an identifier must not be a digit. Identifiers are case-sensitive.

## Numbers

A number like `42` is a non-empty sequence of digits. It may start with a `-` to indicate negative values. Numbers are decimal by default but the base can be overridden with a `0x` prefix for hexadecimal, `0o` for octal and `0b` for binary. The uppercase letters from `A` to `F` are used to represent digits from 10 to 15 in hexadecimal.

## Runes

A rune literal like `'日'` or `'本'` is an integer representing a Unicode code point. It must be enclosed by `'`. It is equivalent to an integer from the perspective of the compiler. The value of the integer is derived from the Unicode representation of the content.

```q
assert 'A' == 0x41
assert 'a' == 0x61
assert '世' == 0x4E16
assert '界' == 0x754C
assert '😀' == 0x1F600
```

## Strings

A string literal like `"Hello"` is a sequence of bytes enclosed by `"`. Strings are immutable, though the compiler does not enforce this rule in its present state. The following escape sequences starting with `\` can be used in rune and string literals to embed special characters:

```q
assert '\0' == 0
assert '\t' == 9
assert '\n' == 10
assert '\r' == 13
assert '\"' == 34
assert '\'' == 39
assert '\\' == 92
```

## Comments

A line comment like `// This is a comment` starts with `//` and stops at the end of the line. Comments are ignored by the compiler and can be added to the code for documentation purposes. Multiline comments are not supported.

## Operators

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

## Keywords

| Keyword  | Description                                                | Stability       |
| -------- | ---------------------------------------------------------- | --------------- |
| `assert` | Tests conditions at runtime                                | ✔️ Stable       |
| `const`  | Defines constant expressions                               | ✔️ Stable       |
| `else`   | Failure branch for if statements                           | ✔️ Stable       |
| `extern` | Foreign function definitions                               | ✔️ Stable       |
| `global` | Global variables (discouraged but required in stdlib)      | ✔️ Stable       |
| `go`     | Asynchronous function calls                                | 🚧 Experimental |
| `if`     | Branches based on a condition                              | ✔️ Stable       |
| `import` | Allows access to other packages                            | ✔️ Stable       |
| `local`  | Thread-local variables                                     | 🚧 Experimental |
| `loop`   | Repeatable code                                            | ✔️ Stable       |
| `return` | Ends the function and returns values to the caller         | ✔️ Stable       |
| `switch` | Multiple branches executing the first true condition block | ✔️ Stable       |

## Builtins

| Function  | Description             | Stability       |
| --------- | ----------------------- | --------------- |
| `cas`     | Atomic compare and swap | 🚧 Experimental |
| `delete`  | Frees memory            | ✔️ Stable       |
| `new`     | Allocates memory        | 🚧 Experimental |
| `syscall` | Calls a kernel function | ✔️ Stable       |

## Packages

A package is defined by a directory.
All files in that directory belong to the same package and share access to its identifiers.
Subdirectories form separate packages.

Within each package, a function named `init` is executed automatically at program startup,
while a function named `exit` runs before the program terminates.
This feature is intended primarily for the standard library and is generally discouraged in application code.

## Library

| Function            | Description                                                          |
| ------------------- | -------------------------------------------------------------------- |
| `bits.rotateLeft`   | Rotate bits left                                                     |
| `c.length`          | Compute the length of a 0-terminated string (C style)                |
| `c.string`          | Create a 0-terminated string (C style)                               |
| `cli.args`          | Slice of arguments passed to the program (excluding executable path) |
| `cli.env`           | Look up environment variables                                        |
| `cli.isTerminal`    | Check if the file handle is a terminal                               |
| `fs.readFile`       | Read a file                                                          |
| `fs.writeFile`      | Write a file                                                         |
| `io.read`           | Read from stdin                                                      |
| `io.readFrom`       | Read from a file handle                                              |
| `io.write`          | Write to stdout                                                      |
| `io.writeTo`        | Write to a file handle                                               |
| `math.newRandom`    | Create a new random number generator                                 |
| `math.sqrt`         | Calculate the square root                                            |
| `mem.alloc`         | Allocate a slice of bytes                                            |
| `mem.copy`          | Copy a slice of bytes                                                |
| `mem.free`          | Free a slice of bytes                                                |
| `mem.zero`          | Zero a slice of bytes                                                |
| `process.id`        | Return the process identifier                                        |
| `process.run`       | Runs the given command as a new process                              |
| `run.exit`          | Exit with an exit code                                               |
| `strings.cut`       | Cut a string in two at a separator                                   |
| `strings.fromInt`   | Convert an integer to a string                                       |
| `strings.index`     | Find a substring                                                     |
| `strings.toInt`     | Convert a string to an integer (base 10)                             |
| `strings.trim`      | Trim whitespace left and right                                       |
| `strings.trimLeft`  | Trim whitespace left                                                 |
| `strings.trimRight` | Trim whitespace right                                                |
| `thread.id`         | Return the thread identifier                                         |
| `time.now`          | Read the current timestamp                                           |
| `time.since`        | Calculate the time passed                                            |
| `time.sleep`        | Sleep the current thread for the given amount of time                |

# Resources

Resources are shared objects such as files, memory or network sockets. The use of resource types prevents the following problems:

- **Resource leaks** (forgetting to free a resource)
- **Use-after-free** (using a resource after it was freed)
- **Double-free** (freeing a resource twice)

Any type, even integers, can be turned into a resource by prefixing the type with `!`. For example, consider these minimal functions:

```q
alloc() -> !int { return 1 }
use(_ int) {}
free(_ !int) {}
```

With this, forgetting to call `free` becomes impossible:

```q
x := alloc()
use(x)
```

```q
x := alloc()
     ┬
     ╰─ Resource of type '!int' not consumed
```

Attempting a use-after-free is also rejected:

```q
x := alloc()
free(x)
use(x)
```

```q
use(x)
    ┬
    ╰─ Unknown identifier 'x'
```

Likewise, a double-free is disallowed:

```q
x := alloc()
free(x)
free(x)
```

```q
free(x)
free(x)
     ┬
     ╰─ Unknown identifier 'x'
```

The compiler only accepts the correct usage order:

```q
x := alloc()
use(x)
free(x)
```

The `!` prefix marks a type to be consumed exactly once. It has no runtime overhead. When a `!int` is passed to another `!int`, the original variable is invalidated in subsequent code. As an exception, converting `!int` to `int` bypasses this rule, allowing multiple uses.

The standard library currently makes use of this feature in two packages:

- `fs.open` must be followed by `fs.close`
- `mem.alloc` must be followed by `mem.free`

For memory allocations of slices and pointers `delete` is called automatically on all exit points of the identifier's scope.
Non-pointer types like `!int` currently do not support automatic life cycle management and require an explicit free call, e.g. an `fs.close` for `!int` file handles.

# Errors

Any function can define an `error` type return value at the end:

```q
a, b, err := canFail()
```

An error value protects all the return values to the left of it.
The protected values `a` and `b` can not be accessed without checking `err` first.
Additionally, error variables like `err` are invalidated after the branch that checked them.

```q
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

# Security

Recent incidents such as the `xz` backdoor and attacks on the `npm` ecosystem have shown that supply chain attacks remain one of the biggest problems in the software industry.

The build system helps mitigate these risks by enforcing the principle of least privilege. Every module must explicitly declare which resources it requires, such as network or file system access. Any permission change becomes part of the review process during updates, making supply chain attacks much more visible. If `leftpad` suddenly requests access to `net`, that should immediately raise suspicion.

While this cannot eliminate supply chain attacks entirely, it significantly reduces the chances of being compromised.

The compiler also hardens executables at the binary level:

- All executables are built as position-independent executables (PIE) with dynamic base addresses so that an attacker can't use precalculated addresses.
- The call stack where return addresses are located is isolated from the regular memory stack, eliminating an entire class of control-flow attacks.
- The W^X (write xor execute) policy is enforced for all memory pages: memory can be writable or executable, but never both.

# Syntax

Q encourages code editors to implement multiple syntaxes for editing.

A view of the code can be substantially different from the underlying model that is saved to disk.
It's important to conceptually realize that one is just a temporary view for editing and the other is a form of persistent data storage.

It is absolutely possible that an editor could offer editing in a Python-like whitespace-significant view.
It is also possible to offer visual editing with a node-based system similar to Scratch or Unreal Engine blueprints.
In all cases the code that is saved to disk would still use the standard text-based format.