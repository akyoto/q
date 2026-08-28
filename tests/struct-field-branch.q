Point {
	x int
	y int
}

global {
	g0 int
	g1 int
}

main() {
	g0 = 2
	g1 = 3
	p := Point{x: 1, y: 2}

	loop i := 0..2 {
		if i > 0 {
			p.y = p.x + g0
		} else {
			g1 = g0 + i
			p.y = p.x + g0
			f(g1)
		}
	}

	assert p.y == 5
}

f(v int) {
	g0 += v
}