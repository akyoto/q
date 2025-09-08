main() {
	x := f
	assert x() == 42
}

f() -> int {
	return 42
}