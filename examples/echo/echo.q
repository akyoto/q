import io

main() {
	buffer := new(byte, 4096)

	loop {
		n, _ := io.read(buffer)

		if n == 0 {
			return
		}

		io.write(buffer[..n])
	}
}