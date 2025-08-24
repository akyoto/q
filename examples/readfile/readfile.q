import fs
import io

main() {
	source := fs.readFile("examples/readfile/readfile.q")
	io.write(source)
}