import io

main() {
	loop i := 0..10 {
		if i % 2 != 0 {
			loop.next()
		}

		io.write(i)
		io.write("\n")
	}
}