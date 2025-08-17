main() {
	a, b := reverse2(1, 2)
	assert a == 2
	assert b == 1
}

reverse2(x int, y int) -> (int, int) {
	return y, x
}