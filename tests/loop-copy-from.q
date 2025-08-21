main() {
	assert f(10) == 20
}

f(n int) -> int {
	b := 0

	loop 0..n {
		b += 2
	}

	return b
}