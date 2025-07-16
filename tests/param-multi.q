main() {
	assert f(1, 2, 3) == 21
}

f(x int, y int, z int) -> int {
	w := g(4, 5, 6)
	return x + y + z + w
}

g(x int, y int, z int) -> int {
	return x + y + z
}