import fs
import io

main() {
	source, err := fs.readFile("examples/readfile/readfile.q")

	if err != 0 {
		io.write("error reading file\n")
		return
	}

	io.write(source)
	delete(source)
}