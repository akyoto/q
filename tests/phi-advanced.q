main() {
	x := 0

	if x == 1 {
		if x == 0 {
			x = 2
		} else {
			x = 3
		}
	} else {
		if x == 0 {
			x = 4
		} else {
			x = 5
		}
	}

	assert x == 4
}