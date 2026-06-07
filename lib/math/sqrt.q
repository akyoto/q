sqrt(n uint) -> uint {
	if n == 0 {
		return 0
	}

	x := n
	y := (x + n / x) >> 1

	loop {
		if y >= x {
			return x
		}

		x = y
		y = (x + n / x) >> 1
	}
}