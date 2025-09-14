Point {
	x int
	y int
}

main() {
	p := Point{x: 1, y: 2}

	if true {
		p = Point{x: 3, y: 4}
	}

	assert p.x == 3
	assert p.y == 4
}