main() {
	c := new(Circle)
	c.center = new(Point)
	c.center.x = 1
	c.center.y = 2
	c.radius = 3

	assert c.center.x == 1
	assert c.center.y == 2
	assert c.radius == 3
}

Circle {
	center *Point
	radius int
}

Point {
	x int
	y int
}