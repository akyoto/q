Point {
	x int
	y int
}

main() {
	p := new(Point){}
	assert p.x == 0
	assert p.y == 0
	delete(p)

	p := new(Point){x: 1}
	assert p.x == 1
	assert p.y == 0
	delete(p)

	p := new(Point){y: 2}
	assert p.x == 0
	assert p.y == 2
	delete(p)

	p := new(Point){x: 1, y: 2}
	assert p.x == 1
	assert p.y == 2
	delete(p)

	p := new(Point){x: 2, y: 2}
	assert p.x == 2
	assert p.y == 2
	delete(p)

	p := new(Point){x: 3, y: 2}
	assert p.x == 3
	assert p.y == 2
	delete(p)

	p := new(Point){x: 3, y: 3}
	assert p.x == 3
	assert p.y == 3
	delete(p)
}