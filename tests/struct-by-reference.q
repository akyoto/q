Point {
	x int
	y int
}

main() {
	p := new(Point)
	assert p.x == 0
	assert p.y == 0
	assert p.x == p.y

	p.x = 1
	p.y = 2
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
	delete(p)
}