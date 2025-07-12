import os

main() {
	os.exit(swap(1, 2))
}

swap(a int, b int) -> int {
	return sum(b, a)
}

sum(a int, b int) -> int {
	return a + b
}