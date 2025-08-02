main() {
	n := 10
	x := 0

	loop {
		if n == 0 {
			assert x == 10
			return
		}

		x = x + 1
		n = n - 1
	}
}