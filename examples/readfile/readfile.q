import cli
import fs
import io

main() {
	args := cli.args()[1..]

	if args.len == 0 {
		io.writeLine("no file specified")
		return
	}

	path := args[0]
	source, err := fs.readFile(path)

	if err != 0 {
		io.write("error reading file: ")
		io.writeLine(path)
		return
	}

	io.write(source)
	delete(source)
}