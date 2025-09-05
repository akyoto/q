import exec
import io
import mem
import strings

main() {
	buffer := mem.alloc(256)

	loop {
		io.write("λ ")
		n := io.read(buffer)

		if n <= 0 {
			mem.free(buffer)
			return
		}

		input := buffer[..n]
		path := strings.trim(input)
		execute(path)
	}
}

execute(path string) {
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