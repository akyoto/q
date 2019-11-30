main() {
	n = 100000

	mut a = 0
	mut b = 0
	mut c = 1

	for 0..n {
		a = b
		b = c
		c = a + b
	}
}
