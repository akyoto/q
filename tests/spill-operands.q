import io

Point {
	x int
	y int
}

main() {
	v0 := -15
	v1 := 19
	v2 := 29
	v3 := -38
	total := 0
	total += rsum(0)
	p1, p2 := pair(v2, v1)
	total += p1 - p2
	s := Point{x: 1, y: 3}
	s.x += 3
	s.y = s.y * 2
	total += s.x + s.y
	pa := new(Point, 3)
	pa[0] = mk(p1, v3)
	pa[1] = mk(v0, s.y)
	pa[2] = mk(p2, v2)
	total += pa[0].x - pa[0].y

	switch {
		s.x > 0 { sel := 11 }
		_       { sel := -4 }
	}

	total += sel
	io.write(total)
}

rsum(n int) -> int {
	if n <= 0 {
		return 3
	}

	return n + rsum(n - 1)
}

pair(x int, y int) -> (int, int) {
	return y + 0, x - 0
}

mk(x int, y int) -> Point {
	return Point{x: x, y: y + 1}
}