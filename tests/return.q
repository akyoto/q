main() {
	assert f(2) == 6
}

f(x int) -> int {
	return x + 1 + g(x)
}

g(x int) -> int {
	return x + 1
}