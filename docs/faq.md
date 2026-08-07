# FAQ

## Which platforms are supported?

|            | arm64 | x86-64 |
| ---------- | ----- | ------ |
| 🐧 Linux   | ✔️    | ✔️     |
| 🍏 Mac     | ✔️    | ✔️     |
| 🪟 Windows | ✔️    | ✔️     |

## How tiny is a Hello World?

|            | arm64    | x86-64  |
| ---------- | -------: | ------: |
| 🐧 Linux   |  0.7 KiB | 0.7 KiB |
| 🍏 Mac     | 32.3 KiB | 8.2 KiB |
| 🪟 Windows |  1.7 KiB | 1.7 KiB |

## Are there any runtime benchmarks?

### Recursive Fibonacci benchmark (`n = 35`):

|                   | arm64                | x86-64               |
| ----------------- | -------------------: | -------------------: |
| C (-O3, gcc 15)   | **41.4 ms** ± 1.4 ms | **24.5 ms** ± 3.2 ms |
| Q (2025-08-20)    | **54.2 ms** ± 1.6 ms | **34.8 ms** ± 2.3 ms |
| Go (1.25, new GC) | **57.7 ms** ± 1.4 ms | **37.9 ms** ± 6.9 ms |
| C (-O0, gcc 15)   | **66.4 ms** ± 1.5 ms | **47.8 ms** ± 4.4 ms |

While the current results lag behind optimized C, this is an expected stage of development. I am actively working to improve the compiler's code generation to a level that can rival optimized C, and I expect a significant performance boost as this work progresses.

### Rewrite of `fnl` in Q:

A rewrite of [fnl](https://git.urbach.dev/cli/fnl) from Rust to Q resulted in a smaller executable with better performance characteristics.

|                            | Lines of code | Binary size | Syscalls | CPU cycles | Time                    |
| -------------------------- | ------------: | ----------: | -------: | ---------: | ----------------------: |
| `fnl` in Q    - 2026-07-08 | 108           |     3.8 KiB | 22       | ~8k        | **270.2 µs** ± 171.6 µs |
| `fnl` in Rust - 2026-06-26 | 111           |   477.2 KiB | 85       | ~500k      | **416.3 µs** ± 224.9 µs |

Note that this is not a 100% fair comparison as the implementation details differ.
Do not trust benchmark data without validating the results yourself.

## Are there any compiler benchmarks?

The table below shows latency numbers on a 2015 Macbook:

|                 | x86-64                  |
| --------------- | ----------------------: |
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
| --------------- | ----------------------: |
| q               |   **89.5 ms** ±  2.4 ms |
| go @1.25        |  **372.2 ms** ±  5.3 ms |
| clang @17.0.0   |  **550.8 ms** ±  3.8 ms |
| rustc @1.89.0   | **1101.0 ms** ±  4.0 ms |
| v @0.4.11       | **1256.0 ms** ±  4.0 ms |
| zig @0.15.1     | **1407.0 ms** ± 12.0 ms |
| odin @accdd7c2a | **1770.0 ms** ±  7.0 ms |

## What is the compiler based on?

The backend is built on a [Static Single Assignment (SSA)](https://en.wikipedia.org/wiki/Static_single-assignment_form) intermediate representation, the same approach used by mature compilers such as `gcc`, `go`, and `llvm`. SSA greatly simplifies the implementation of common optimization passes, allowing the compiler to produce relatively high-quality assembly code despite the project's early stage of development.

## Any editor extensions?

- **Neovim**: Planned.
- **VS Code**: Clone the [vscode-q](https://git.urbach.dev/extra/vscode-q) repository into your extensions folder (it enables syntax highlighting).
- **Zed**: Planned.

## Can I use it for scripting?

Yes. The compiler can run an entire script within a few microseconds.

```q
!/usr/bin/env q
import io

main() {
	io.write("Hello\n")
}
```

You need to create a file with the contents above and add execution permissions via `chmod +x`. Now you can run the script without an explicit compiler build. The generated machine code runs directly from RAM if the OS supports it.

## What's the minimum supported version of...?

|                          | Version       | Year | Reason                |
| ------------------------ | ------------- | ---- | --------------------- |
| 🐧 Linux                 | 5.16          | 2022 | futex2                |
| 🍏 Mac                   | 11.0          | 2020 | Chained fixups        |
| 🪟 Windows               | 8.1           | 2013 | FSGSBASE              |
| 🔲 x86-64 (Intel)        | Ivy Bridge    | 2012 | FSGSBASE              |
| 🔲 x86-64 (AMD)          | Excavator     | 2015 | FSGSBASE              |
| 🔲 arm64                 | 8.1-A         | 2014 | CAS                   |
| 💻 Raspberry Pi          | 5             | 2023 | CAS                   |

If you need support for older hardware and software, feel free to send a PR with the required fallback mechanisms.

## Why is it written in Go and not language X?

Because of readability and great tools for concurrency.
The implementation will be replaced by a self-hosted compiler in the future.

## How do I pronounce the name?

/ˈkjuː/ just like `Q` in the English alphabet.
