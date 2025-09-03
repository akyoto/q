main() {
	a, b, _ := f()
	assert a == 1
	assert b == 2
}

f() -> (int, int, error) {
	return 1, 2, 3
}