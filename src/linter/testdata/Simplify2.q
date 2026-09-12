main() {
	x := 42

	loop 0..1 {
		x = x | 0
	}

	assert x == 42
}