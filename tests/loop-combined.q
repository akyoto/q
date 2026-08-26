import io

main() {
	loop i := 0..10 {
		if i == 2 {
			i = 5
		}

		if i == 5 {
			loop.next()
		}

		if i == 9 {
			loop.stop()
		}

		io.write(i)
	}
}