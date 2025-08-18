Point {
	x int
	y int
}

main() {
	p := Point{x: 1, y: 2}
	assert p.x == 1
	assert p.y == 2
	assert p.x != p.y

	p.x = p.y
	assert p.x == 2
	assert p.y == 2
	assert p.x == p.y

	p.x = p.y + 1
	assert p.x == 3
	assert p.y == 2
	assert p.x != p.y

	p.y += 1
	assert p.x == 3
	assert p.y == 3
	assert p.x == p.y
}