main() {
	x := 1

	loop i := 0..4 {
		x += i
	}

	assert x == 7
}