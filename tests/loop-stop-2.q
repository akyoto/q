main() {
	n := 0

	loop {
		if n == 5 {
			loop.stop()
		}

		if n == 6 {
			n = 7
			loop.stop()
		}

		assert n < 5
		n += 1
	}

	assert n == 5
}