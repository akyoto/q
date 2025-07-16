main() {
	assert f(1) == 3
}

f(x int) -> int {
	y := g()
	return x + y
}

g() -> int {
	return 2
}