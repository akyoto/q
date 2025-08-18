Point {
	x int
	y int
}

Point(x int, y int) -> *Point {
	p := new(Point)
	p.x = x
	p.y = y
	return p
}

main() {
	a := Point(1, 2)
	assert a.x == 1
	assert a.y == 2

	b := Point(3, 4)
	assert b.x == 3
	assert b.y == 4
}