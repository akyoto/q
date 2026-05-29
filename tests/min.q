main() {
	a := 1
	b := 2
	min := a

	if b < min {
		min = b
	}

	assert min == 1
}