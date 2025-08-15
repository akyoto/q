Point {
	x int
	y int
}

main() {
	p := new(Point)
	assert p.x == 0
	assert p.y == 0

	p.x = 1
	p.y = 2
	assert p.x == 1
	assert p.y == 2
}