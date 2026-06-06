main() {
	x, y, z := f()
	assert x == 1
	assert y == 2
	assert z == 3
}

f() -> (int, int, int) {
	return g()
}

g() -> (int, int, int) {
	return 1, 2, 3
}