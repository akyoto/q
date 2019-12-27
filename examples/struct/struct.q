import sys

struct Point {
	x Int
	y Int
}

main() {
	let p1 = Point()
	p1.x = 10
	p1.y = 20

	let p2 = Point()
	p2.x = 30
	p2.y = p1.y

	let s = sum(p1, p2)
	sys.exit(s)
}

sum(a Point, b Point) -> Int {
	return a.x + b.x + a.y + b.y
}
