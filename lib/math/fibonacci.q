fibonacci(n) {
	require n >= 0
	ensure _ >= 0

	mut a = ?
	mut b = 0
	mut c = 1

	for 0..n {
		a = b
		b = c
		c = a + b
	}

	return b
}
