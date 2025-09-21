main() {
	p := new(Point, 2)
	assert p.len == 2

	p[0].x = 1
	p[0].y = 2
	p[1].x = 3
	p[1].y = 4

	assert p[0].x == 1
	assert p[0].y == 2
	assert p[1].x == 3
	assert p[1].y == 4
}

Point {
	x int
	y int
}