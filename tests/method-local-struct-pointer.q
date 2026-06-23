import io

main() {
	p := new(Point){x: 1, y: 2}
	p.print()
}

Point {
	x int
	y int
}

print(p *Point) {
	io.writeLine(p.x)
	io.writeLine(p.y)
}