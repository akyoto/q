import io

Point {
	x int
	y int
}

main() {
	p := Point(1, 2)
	write(p)
}

Point(x int, y int) -> *Point {
	p := new(Point)
	p.x = x
	p.y = y
	return p
}

write(p *Point) {
	io.write("Point: ")
	io.writeInt(p.x)
	io.write(", ")
	io.writeInt(p.y)
}