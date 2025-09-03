import fs
import io
import mem

main() {
	source, err := fs.readFile("examples/readfile/readfile.q")

	if err != 0 {
		io.write("error reading file\n")
		return
	}

	io.write(source)
	mem.free(source)
}