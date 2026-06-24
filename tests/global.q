main() {
	assert x == 0
	f(42)
	assert x == 42
}

f(n int) {
	x = n
}

global {
	x int
}