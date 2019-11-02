# Q

The Q programming language.

This is a very early version of a programming language I'm currently working on.
Don't expect it to do anything useful yet, it can barely create a Linux executable printing "Hello World".

## Features

* Simple (inspired by Go)
* Fast compilation (inspired by V)
* Memory-safe (inspired by Rust)
* Functional (inspired by Haskell)
* Zero-cost abstractions (inspired by Rust)
* No dependencies (no LLVM, no GCC, ...)

## Installation

```shell
git clone https://github.com/akyoto/q
go build
./q testdata/hello.q
./hello
```

## Design goals

* Importing a module should have no side effects
* Modules should be highly reusable (not bound to predefined data structures)
* No classes or methods, instead we have data and functions that can operate on data
* No name shadowing, names never change their meaning
* ...TODO...

## Implementation goals

* User-friendly compiler messages
* ...TODO...

## How to contribute

* ...TODO...
