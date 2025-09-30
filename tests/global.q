main() {
	assert x == 0
	f()
	assert x == 42
}

f() {
	x = 42
}

global {
	x int
}