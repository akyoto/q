import cli
import fs
import io
import mem
import run

const {
	MODE_STATUS = 0
	MODE_ADD = 1
	MODE_REMOVE = 2
}

main() {
	mode := MODE_STATUS
	args := cli.args()

	loop i := 0..args.len {
		switch {
			args[i] == "-add" {
				mode = MODE_ADD
			}
			args[i] == "-remove" {
				mode = MODE_REMOVE
			}
			_ {
				processFile(args[i], mode)
			}
		}
	}

	if !cli.isTerminal(io.stdin) {
		processStdin(mode)
	}
}

processStdin(mode int) {
	buffer := new(byte, 0x4000)
	pos := 0

	loop {
		n, _ := io.read(buffer[pos..])

		if n == 0 {
			switch {
				mode == MODE_ADD && pos < buffer.len && buffer[pos-1] != '\n' {
					buffer[pos] = '\n'
					pos += 1
				}
				mode == MODE_REMOVE && pos > 0 && buffer[pos-1] == '\n' {
					pos -= 1
				}
			}

			io.write(buffer[..pos])
			return
		}

		pos += n

		if pos == buffer.len {
			// TODO: This is not fully correct yet.
			io.write(buffer)
			pos = 0
		}
	}
}

processFile(path string, mode int) {
	source, err := fs.readFile(path)

	if err != 0 {
		io.write("error reading file: ")
		io.writeLine(path)
		run.exit(err)
	}

	if source.len == 0 {
		return
	}

	io.write(path)
	io.write(": ")

	if source[source.len-1] == '\n' {
		if mode == MODE_REMOVE {
			io.writeLine("no final newline [removed]")
			fs.writeFile(path, source[..source.len-1])
		} else {
			io.writeLine("final newline")
		}
	} else {
		if mode == MODE_ADD {
			io.writeLine("final newline [added]")
			newSource := addNewline(source)
			fs.writeFile(path, newSource)
			delete(newSource)
		} else {
			io.writeLine("no final newline")
		}
	}
}

addNewline(content string) -> !string {
	newContent := new(byte, content.len+1)
	mem.copy(newContent, content)
	newContent[content.len] = '\n'
	return newContent
}