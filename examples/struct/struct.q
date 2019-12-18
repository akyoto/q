import sys

struct Point {
	x Int
	y Int
}

main() {
	let p = Point()
	p.x = 10
	p.y = 20
	sys.exit(p.y)
}
