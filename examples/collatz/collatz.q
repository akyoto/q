main() {
	collatz(12)
}

collatz(x int) {
	loop {
		if x & 1 == 0 {
			x = x / 2
		} else {
			x = 3 * x + 1
		}

		if x == 1 {
			return
		}
	}
}