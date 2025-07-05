<p align="center"><img alt="q logo" src="logo.svg"></p>

> [!NOTE]
> `q` is under heavy development and not ready for production yet.
>
> Feel free to [get in touch](https://urbach.dev/contact) if you are interested in helping out.

## ⚡️ Installation

```shell
git clone https://git.urbach.dev/cli/q
cd q
go build
```

To symlink the compiler as `q`:

```shell
ln -s $PWD/q ~/.local/bin/q
```

## 🚀 Usage

```shell
q build examples/hello
./examples/hello/hello
```

## 🚦 Tests

```shell
go run gotest.tools/gotestsum@latest
```

## 💻 Platforms

|         | arm64  | x86-64 |
| ------- | ------ | ------ |
| Linux   | ✔️     | ✔️     |
| Mac     | ✔️*    | ✔️     |
| Windows | ✔️*    | ✔️     |

Those marked with a star are supported in theory but there are no developer machines to test them.

## 🔑 Security

### PIE

All executables are built as Position Independent Executables (PIE) supporting a dynamic base address.

### Memory pages

Code and data are separated into different memory pages and loaded with different access permissions.

|        | Read | Execute | Write |
| ------ | ---- | ------- | ----- |
| Code   | ✔️   | ✔️      | ❌    |
| Data   | ✔️   | ❌      | ❌    |

## 🧾 License

Please see the [license documentation](https://urbach.dev/license).

## 🧔 Copyright

© 2025 Eduard Urbach