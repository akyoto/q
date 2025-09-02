import io
import mem
import run

main() {
	buffer := mem.alloc(6)
	fill(buffer)
	check(buffer)

	loop i := 0..buffer.len {
		writeByte(buffer[i])
	}

	mem.free(buffer)
}

fill(buffer string) {
	buffer[0] = 'H'
	buffer[1] = 'e'
	buffer[2] = 'l'
	buffer[3] = 'l'
	buffer[4] = 'o'
	buffer[5] = '\n'
}

check(buffer string) {
	assert buffer[0] == 'H'
	assert buffer[1] == 'e'
	assert buffer[2] == 'l'
	assert buffer[3] == 'l'
	assert buffer[4] == 'o'
	assert buffer[5] == '\n'
}

writeByte(n byte) {
	switch {
		n == 'H'  { io.write("H") }
		n == 'e'  { io.write("e") }
		n == 'l'  { io.write("l") }
		n == 'o'  { io.write("o") }
		n == '\n' { io.write("\n") }
		_         { run.exit(1) }
	}
}