main() {
	a := 1
	b := 2
	max := a

	if b > max {
		max = b
	}

	assert max == 2
}