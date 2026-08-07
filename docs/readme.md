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

- High performance (comparable to C and Go)
- Fast compilation (5-10x faster than most compilers)
- Lightweight executables (1 KB for simple programs)
- Static analysis (integrated linter catches common mistakes)
- Pointer safety (pointers cannot be nil)
- Resource safety (use-after-free is a compile error)
- Simple syntax (control flow is easily understood)
- Friendly errors (clear and concise compiler messages)
- General purpose (apps, servers, games, kernels, etc.)
- Multiple architectures (x86-64 and arm64)
- Multiple platforms (Linux, Mac and Windows)
- Readable source (less than 1% of LLVM's code size)
- Zero dependencies (no external tools or libraries)

## Quickstart

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

- **IRC**: [#q](ircs://irc.urbach.dev:6697/#q) on [irc.urbach.dev](https://irc.urbach.dev)
- **Discord**: [Q community](https://discord.gg/4q3DJFsTvB)

## Donate

- [GitHub](https://github.com/sponsors/akyoto)
- [Kofi](https://ko-fi.com/akyoto)

## License

See the [license documentation](https://urbach.dev/license).

## Copyright

© 2025 Eduard Urbach