import io

main() {
	n := 10

	loop {
		if n < 0 {
			return
		}

		io.writeLine(n)
		n -= 1
	}
}