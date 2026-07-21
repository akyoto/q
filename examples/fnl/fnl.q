import cli
import fs
import io
import mem
import run

Mode const {
	Status = 0
	Add = 1
	Remove = 2
}

main() {
	mode := Mode.Status
	args := cli.args()

	loop i := 0..args.len {
		switch args[i] {
			"-add"    { mode = Mode.Add }
			"-remove" { mode = Mode.Remove }
			_ {
				err := processFile(args[i], mode)

				if err != 0 {
					io.write("error processing file '")
					io.write(args[i])
					io.write("': ")
					io.writeLine(err)
					run.exit(1)
				}

				return
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
				mode == Mode.Add && pos < buffer.len && buffer[pos-1] != '\n' {
					buffer[pos] = '\n'
					pos += 1
				}
				mode == Mode.Remove && pos > 0 && buffer[pos-1] == '\n' {
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

processFile(path string, mode int) -> error {
	source, err := fs.readFile(path)

	if err != 0 {
		return err
	}

	if source.len == 0 {
		return 0
	}

	io.write(path)
	io.write(": ")

	if source[source.len-1] == '\n' {
		if mode == Mode.Remove {
			io.writeLine("no final newline [removed]")
			return fs.writeFile(path, source[..source.len-1])
		} else {
			io.writeLine("final newline")
		}
	} else {
		if mode == Mode.Add {
			io.writeLine("final newline [added]")
			newSource := addNewline(source)
			return fs.writeFile(path, newSource)
		} else {
			io.writeLine("no final newline")
		}
	}

	return 0
}

addNewline(content string) -> !string {
	newContent := new(byte, content.len+1)
	mem.copy(newContent, content)
	newContent[content.len] = '\n'
	return newContent
}