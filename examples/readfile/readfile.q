import fs
import io
import mem

main() {
	source := fs.readFile("examples/readfile/readfile.q")
	io.write(source)
	mem.free(source)
}