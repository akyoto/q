Point {
	x int
	y int
}

main() {
	p := Point{x: 1, y: 2}
	assert p.x == 1
	assert p.y == 2
}