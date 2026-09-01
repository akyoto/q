main() {
	a := 4
	b := 8
	c := 7

	if a != b {
		c = a
	}

	assert c == 4
}