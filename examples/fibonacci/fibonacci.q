import os

main() {
	os.exit(fibonacci(10))
}

fibonacci(n int) -> int {
	if n <= 1 {
		return n
	}

	return fibonacci(n - 1) + fibonacci(n - 2)
}