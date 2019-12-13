import sys

main() {
	# Calculate n'th fibonacci number
	let n = 11

	mut b = 0
	mut c = 1

	for 0..n {
		let a = b
		b = c
		c = a + b
	}

	sys.exit(b)
}
