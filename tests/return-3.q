main() {
	a, b, c := reverse3(1, 2, 3)
	assert a == 3
	assert b == 2
	assert c == 1
}

reverse3(x int, y int, z int) -> (z int, y int, x int) {
	return z, y, x
}