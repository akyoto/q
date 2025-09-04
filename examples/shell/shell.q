import exec
import io
import mem

main() {
	length := 256
	buffer := mem.alloc(length)

	loop {
		io.write("λ ")
		n := io.read(buffer)

		if n <= 0 {
			mem.free(buffer)
			return
		}

		if buffer[n-1] == '\n' {
			n -= 1
		}

		if n > 0 && buffer[n-1] == '\r' {
			n -= 1
		}

		cmd(buffer[..n])
	}
}

cmd(path string) {
	if path.len == 0 {
		return
	}

	err := exec.run(path)

	if err != 0 {
		io.write("error executing: ")
		io.write(path)
		io.write("\n")
	}
}