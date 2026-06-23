main() {
	n := 10

	loop 0..n {
		n -= 1
	}

	assert n == 5
}