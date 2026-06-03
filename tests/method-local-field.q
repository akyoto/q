import io

main() {
	f().x.print()
	f().y.print()
}

print(x int) {
	io.writeLine(x)
}

f() -> Point {
	return Point{x: 1, y: 2}
}

Point {
	x int
	y int
}