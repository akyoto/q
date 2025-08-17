Point {
	x int
	y int
}

main() {
	p := Point{x: 1, y: 2}
	assert p.x == 1
	assert p.y == 2

	p.x = 3
	p.y = 4
	assert p.x == 3
	assert p.y == 4
}