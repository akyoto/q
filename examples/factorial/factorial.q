import os

main() {
	os.exit(factorial(5))
}

factorial(n int) -> int {
	if n <= 1 {
		return 1
	}

	return n * factorial(n - 1)
}