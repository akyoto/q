import io

main() {
	buffer := new(byte, 4096)

	loop {
		n := io.read(buffer)

		if n <= 0 {
			delete(buffer)
			return
		}

		io.write(buffer[..n])
	}
}