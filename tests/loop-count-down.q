main() {
	n := 10
	x := 0

	loop {
		if n == 0 {
			assert x == 10
			return
		}

		x += 1
		n -= 1
	}
}