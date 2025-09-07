main() {
	a, b, c, d := reverse4(1, 2, 3, 4)
	assert a == 4
	assert b == 3
	assert c == 2
	assert d == 1

	a, b, c, d = reverse4(10, 20, 30, 40)
	assert a == 40
	assert b == 30
	assert c == 20
	assert d == 10

	a, _, c, _ = reverse4(100, 200, 300, 400)
	assert a == 400
	assert b == 30
	assert c == 200
	assert d == 10
}

reverse4(x int, y int, z int, w int) -> (int, int, int, int) {
	return w, z, y, x
}