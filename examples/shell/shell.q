import io
import process
import strings

main() {
	buffer := new(byte, 256)

	loop {
		io.write("λ ")
		n, _ := io.read(buffer)

		if n == 0 {
			delete(buffer)
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

	err := process.run(path)

	if err != 0 {
		io.write("error executing: ")
		io.writeLine(path)
	}
}