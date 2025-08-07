import io

main() {
	collatz(12)
}

collatz(x int) {
	loop {
		io.writeInt(x)

		if x == 1 {
			return
		}

		io.write(" ")

		if x & 1 == 0 {
			x = x / 2
		} else {
			x = 3 * x + 1
		}
	}
}