main() {
	c := Circle{
		center: Point{x: 1, y: 2},
		radius: 3
	}

	assert c.center.x == 1
	assert c.center.y == 2
	assert c.radius == 3
}

Circle {
	center Point
	radius int
}

Point {
	x int
	y int
}