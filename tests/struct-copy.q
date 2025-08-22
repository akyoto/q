Point {
	x int
	y int
}

main() {
	a := Point{x: 1, y: 2}
	b := a
	assert a.x == b.x
	assert a.y == b.y

	a.x = 3
	a.y = 4
	assert a.x != b.x
	assert a.y != b.y

	a = b
	assert a.x == b.x
	assert a.y == b.y
}