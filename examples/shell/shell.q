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

		cmd(buffer[..n-1])
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