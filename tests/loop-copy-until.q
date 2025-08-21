main() {
	a := 0
	b := 10
	n := 10

	loop 0..n {
		a += 1
		b -= 1
	}

	assert a == 10
	assert b == 0
}