main() {
	assert f(1) == 3
}

f(x int) -> int {
	return x + 1 + x
}