import io

Point {
	x int
	y int
}

main() {
	p := new(Point)
	p.x = 1
	p.y = 2

	io.write("Point: ")
	io.writeInt(p.x)
	io.write(", ")
	io.writeInt(p.y)
}