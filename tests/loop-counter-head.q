import io

main() {
	n := 10

	loop {
		n -= 1

		if n < 0 {
			return
		}

		io.writeLine(n)
	}
}