Point {
	x int
	y int
}

Point(x int, y int) -> Point {
	return Point{x: x, y: y}
}

main() {
	a := Point(1, 2)
	assert a.x == 1
	assert a.y == 2

	b := Point(3, 4)
	assert b.x == 3
	assert b.y == 4
}