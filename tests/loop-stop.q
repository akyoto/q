import io

main() {
	n := 0

	loop {
		if n == 5 {
			loop.stop()
		}

		io.write(n)
		n += 1
	}

	io.write("\n")
	io.write(n)
}