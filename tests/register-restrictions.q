main() {
	a := 0
	b := 1

	if true {
		a += 1
		b += 1
	}

	assert a == 1
	assert b == 2
}