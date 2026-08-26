main() {
	x := 0

	loop i := 0..4 {
		if i == 2 {
			loop.next()
		}

		x += i
	}

	assert x == 4
}