import io

Point {
	x int
	y int
}

Point(x int, y int) -> *Point {
	p := new(Point)
	p.x = x
	p.y = y
	return p
}

main() {
	p := Point(1, 2)
	write(p)
}

write(p *Point) {
	io.write("Point: ")
	io.write(p.x)
	io.write(", ")
	io.write(p.y)
}