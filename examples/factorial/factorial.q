import io

main() {
	io.write(factorial(5))
}

factorial(n int) -> int {
	if n <= 1 {
		return 1
	}

	return n * factorial(n - 1)
}