import io

Point {
	x int
	y int
}

main() {
	p := new(Point)
	p.x = 1
	p.y = 2
	write(p)
	delete(p)
}

write(p *Point) {
	io.write("Point: ")
	io.write(p.x)
	io.write(", ")
	io.write(p.y)
}