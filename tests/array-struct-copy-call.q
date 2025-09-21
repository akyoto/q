main() {
	p := new(Point, 2)
	p[0] = f(1, 2)
	p[1] = f(3, 4)
	assert p[0].x == 1
	assert p[0].y == 2
	assert p[1].x == 3
	assert p[1].y == 4
}

f(x int, y int) -> Point {
	return Point{x: x, y: y}
}

Point {
	x int
	y int
}