# Reference

> [!WARNING]
> Q is in early development and not ready for production yet.
>
> The reference manual is work in progress.

Q is a programming language for building correct and high-performance systems software.

## Tokens

Source files are preprocessed by the tokenizer which groups the individual bytes into several token types:

- Invalid
- EOF
- NewLine
- [Identifier](#identifiers)
- Number
- Rune
- String
- Comment
- Script
- GroupStart / GroupEnd
- BlockStart / BlockEnd
- ArrayStart / ArrayEnd
- ReturnType
- [Operators](#operators)
- [Keywords](#keywords)
- [Builtins](#builtins)

## Identifiers

An identifier is a non-empty sequence of letters, digits, and underscores (`_`).
The first character of an identifier must not be a digit. Identifiers are case-sensitive.

## Operators

Operator precedence introduces additional rules that programmers must learn and can lead to hidden mistakes.
To minimize this complexity, the language uses only 9 precedence levels:

| Precedence | Operators                                                         | Description             |
| ---------: | ----------------------------------------------------------------- | ----------------------- |
|          9 | `.` `()` `[]` `{}`                                                | Postfix                 |
|          8 | `!` `-`                                                           | Unary                   |
|          7 | `*` `/` `%`                                                       | Multiplicative          |
|          6 | `+` `-` `&` `\|` `^` `<<` `>>` `as`                               | Additive, bitwise, cast |
|          5 | `==` `!=` `<` `>` `<=` `>=`                                       | Comparison              |
|          4 | `&&`                                                              | Logical AND             |
|          3 | `\|\|`                                                            | Logical OR              |
|          2 | `..` `,`                                                          | Range, separator        |
|          1 | `:=` `=` `+=` `-=` `*=` `/=` `%=` `&=` `\|=` `^=` `<<=` `>>=` `:` | Assignment              |

## Keywords

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

## Builtins

| Function   | Description                                                | Stability       |
| ---------- | ---------------------------------------------------------- | --------------- |
| `delete`   | Frees memory                                               | 🚧 Experimental |
| `new`      | Allocates memory                                           | 🚧 Experimental |
| `syscall`  | Calls a kernel function                                    | ✔️ Stable       |

## Packages

A package is defined by a directory.
All files in that directory belong to the same package and share access to its identifiers.
Subdirectories form separate packages.

Within each package, a function named `init` is executed automatically at program startup,
while a function named `exit` runs before the program terminates.
This feature is intended primarily for the standard library and is generally discouraged in application code.
