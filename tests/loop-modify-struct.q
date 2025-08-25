import io

Point {
	x int
	y int
}

main() {
	p := Point{x: 1}

	loop 0..5 {
		io.write(p.x)
		p.x += 1
	}
}