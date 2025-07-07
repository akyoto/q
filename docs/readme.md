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
* Fast compilation (5-10x faster than most)
* Unix scripting (JIT compilation)
* No dependencies (no llvm, no libc)

## Installation

```shell
git clone https://git.urbach.dev/cli/q
cd q
go build
```

## Usage

Quick test:

```shell
q run examples/hello
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

The compiler is actually so fast that it's possible to use `q` for scripting. Create a new file:

```q
#!/usr/bin/env q

import io

main() {
	io.write("Hello\n")
}
```

Add permissions via `chmod +x`. The file can be executed from anywhere now.
The machine code is run directly from memory if the OS supports it.

## Tests

```shell
go run gotest.tools/gotestsum@latest
```

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