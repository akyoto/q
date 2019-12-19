import sys

struct Point {
	x Int
	y Int
}

main() {
	let p = Point()
	p.x = 20
	p.y = p.x
	sys.exit(p.y)
}
