<div align="center">
	<img src="logo.svg" width="90" alt="Q logo">
	<h1>The Q Programming Language</h1>
	<p>
		<a href="#features">Features</a> ·
		<a href="#quickstart">Quickstart</a> ·
		<a href="#examples">Examples</a> ·
		<a href="#docs">Docs</a>
	</p>
	<p>
		A minimal, dependency-free language that compiles to tiny, fast machine code.
	</p>
</div>

## Features

- 🚄 **High performance** - <small>comparable to Rust and Zig.</small>
- 🚀 **Fast compilation** - <small>5-10x faster than most compilers.</small>
- 📦 **Lightweight executables** - <small>1 KB for simple programs.</small>
- 🔍 **Static analysis** - <small>integrated linter catches common mistakes.</small>
- 🛡️ **Pointer safety** - <small>pointers cannot be nil.</small>
- ♻️ **Resource safety** - <small>use-after-free is a compile error.</small>
- 🧠 **Simple syntax** - <small>control flow is easily understood.</small>
- 💬 **Friendly errors** - <small>clear and concise compiler messages.</small>
- 🌐 **General purpose** - <small>apps, servers, games, kernels, etc.</small>
- 🧩 **Multiple architectures** - <small>x86-64 and arm64.</small>
- 🖥️ **Multiple platforms** - <small>Linux, Mac and Windows.</small>
- 📖 **Readable source** - <small>less than 1% of LLVM's code size.</small>
- 🧘 **Zero dependencies** - <small>no external tools or libraries.</small>

## Quickstart

> [!WARNING]
> Q is in early development. Design and implementation are changing continuously.

### Install

```shell
git clone https://git.urbach.dev/cli/q
cd q
go build
ln -s $PWD/q ~/.local/bin/
```

### Run

```shell
q examples/hello
```

### Build

```shell
q build examples/hello
```

Cross-compile with `-os`:

```shell
q build examples/hello -os windows
```

### Inspect

```shell
q ssa examples/hello   # SSA form
q asm examples/hello   # Assembly output
```

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

## Docs

- [Motivation](motivation.md)
- [Design](design.md)
- [Reference](reference.md)
- [Source](source.md)
- [Changes](changes.md)
- [Security](security.md)
- [Contributing](contributing.md)
- [FAQ](faq.md)

## Community

- [IRC](https://irc.urbach.dev/#q)
- [Discord](https://discord.gg/4q3DJFsTvB)

## Donate

- [GitHub](https://github.com/sponsors/akyoto)
- [Kofi](https://ko-fi.com/akyoto)

## License

See the [license documentation](https://urbach.dev/license).

## Copyright

© 2025 Eduard Urbach