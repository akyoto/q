import cli
import fs
import io

main() {
	args := cli.args()

	if args.len == 0 {
		io.writeLine("no files")
		return
	}

	loop i := 0..args.len {
		read(args[i])
	}
}

read(path string) {
	source, err := fs.readFile(path)

	if err != 0 {
		io.write("error reading file: ")
		io.writeLine(path)
		return
	}

	io.writeLine(source)
}