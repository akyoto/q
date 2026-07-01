import io

main() {
	loop i := 1..15 {
		io.writeLine(factorial(i))
	}
}

factorial(n uint) -> uint {
	if n <= 1 {
		return 1
	}

	return n * factorial(n - 1)
}