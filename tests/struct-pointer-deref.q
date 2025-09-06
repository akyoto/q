main() {
	p := new(Point)
	p.x = 1
	p.y = 2

	p2 := [p]
	p2.x = 3

	assert p.x == 1
	assert p.y == 2
	assert p2.x == 3
	assert p2.y == 2

	delete(p)
}

Point {
	x int
	y int
}