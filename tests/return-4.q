main() {
	a, b, c, d := mix4(1, 2, 3, 4)
	assert a == 4 + 1
	assert b == 3 + 2
	assert c == 2 + 3
	assert d == 1 + 4
}

mix4(x int, y int, z int, w int) -> (int, int, int, int) {
	return w + x, z + y, y + z, x + w
}