import process

main() {
	n = 11

	mut a = 0
	mut b = 0
	mut c = 1

	for 1..n {
		a = b
		b = c
		c = a + b
	}

	process.exit(c)
}
