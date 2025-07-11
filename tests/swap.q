import os

main() {
	os.exit(f(1, 2))
}

f(a int, b int) -> int {
	return sum(b, a)
}

sum(a int, b int) -> int {
	return a + b
}