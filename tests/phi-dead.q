main() {
	assert f(-42) == -42
}

f(b int) -> int {
	a := 1

	if b >= 0 {
		// ...
	} else {
		if a <= 1 {
			a = 0
		} else {
			b = -21
		}
	}

	return b
}