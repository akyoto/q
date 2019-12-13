fibonacci(n Int) -> Int {
	expect n >= 0
	ensure _ >= 0

	mut b = 0
	mut c = 1

	for 0..n {
		let a = b
		b = c
		c = a + b
	}

	return b
}
