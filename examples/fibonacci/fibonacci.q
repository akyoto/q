import io

main() {
	loop i := 1..15 {
		io.writeLine(fibonacci(i))
	}
}

fibonacci(n uint) -> uint {
	if n <= 1 {
		return n
	}

	return fibonacci(n - 1) + fibonacci(n - 2)
}