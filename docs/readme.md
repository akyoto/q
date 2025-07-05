# q

A programming language including compiler, assembler and linker with the following goals:

- Fast compilation
- High performance
- Tiny executables

## Installation

Build:

```shell
git clone https://git.urbach.dev/cli/q
cd q
go build
```

Symlink:

```shell
ln -s $PWD/q ~/.local/bin/q
```

## Usage

Build:

```shell
q build examples/hello
```

Run:

```shell
./examples/hello/hello
```

## Tests

To run all tests:

```shell
go run gotest.tools/gotestsum@latest
```

Every package must pass its unit tests.

The code coverage should ideally be around 80-100% for each package.

Integration tests that test the output of the produced executables are located in the dedicated `tests` directory.

## Platforms

### Matrix

|         | arm64  | x86-64 |
| ------- | ------ | ------ |
| Linux   | ✔️     | ✔️     |
| Mac     | ✔️*    | ✔️     |
| Windows | ✔️*    | ✔️     |

Those marked with a star are supported in theory but there are no developer machines to test them.

### Cross compilation

Compilation for a different target is supported out of the box:

```shell
q build examples/hello --os linux
q build examples/hello --os mac
q build examples/hello --os windows
```

```shell
q build examples/hello --arch x86
q build examples/hello --arch arm
```

## Security

### PIE

All executables are built as Position Independent Executables (PIE) supporting a dynamic base address.

### Memory pages

Code and data are separated into different memory pages and loaded with different access permissions.

|        | Read | Execute | Write |
| ------ | ---- | ------- | ----- |
| Code   | ✔️   | ✔️      | ❌    |
| Data   | ✔️   | ❌      | ❌    |

## Status

`q` is under heavy development and not ready for production yet.

Feel free to [get in touch](https://urbach.dev/contact) if you are interested in helping out.

## License

Please see the [license documentation](https://urbach.dev/license).

## Copyright

© 2025 Eduard Urbach