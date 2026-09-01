main() {
	x := 1
	y := 2

	z := switch {
		x < y { 10 }
		_     { 20 }
	}

	assert z == 10
}