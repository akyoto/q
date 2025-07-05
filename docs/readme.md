# q

A programming language that quickly compiles to machine code.

## Goals

- Fast compilation
- High performance
- Tiny executables

## Installation

```shell
git clone https://git.urbach.dev/cli/q
cd q
go build
```

## Usage

```shell
q build examples/hello
./examples/hello/hello
```

## Tests

```shell
go run gotest.tools/gotestsum@latest
```

## Platforms

You can cross-compile executables for Linux, Mac and Windows using arm64 or x86-64.

```shell
q build examples/hello --os linux
q build examples/hello --os mac
q build examples/hello --os windows
q build examples/hello --arch x86
q build examples/hello --arch arm
```

| CPU    | Linux | Mac | Windows |
| ------ | ----- | --- | ------- |
| arm64  | ✔️    | ❔  | ❔      |
| x86-64 | ✔️    | ✔️  | ✔️      |

## Status

`q` is under heavy development and not ready for production yet.
Feel free to [get in touch](https://urbach.dev/contact) if you are interested in helping out.

## License

Please see the [license documentation](https://urbach.dev/license).

## Copyright

© 2025 Eduard Urbach