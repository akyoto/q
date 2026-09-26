Grid {
	data [2][3]int
}

main() {
	grid := new(Grid)
	grid.data[1][2] = 42
	a := grid.data[1][1..3]
	assert a[1] == 42
	b := grid.data[1][..3]
	assert b[2] == 42
}