Point {
	x int
	y int
}

main() {
	a := new(Point)
	assert a.x == 0
	assert a.y == 0
	assert a.x == a.y

	a.x = 1
	a.y = 2
	assert a.x == 1
	assert a.y == 2
	assert a.x != a.y

	a.x = a.y
	assert a.x == 2
	assert a.y == 2
	assert a.x == a.y

	a.x = a.y + 1
	assert a.x == 3
	assert a.y == 2
	assert a.x != a.y

	a.y += 1
	assert a.x == 3
	assert a.y == 3
	assert a.x == a.y
}