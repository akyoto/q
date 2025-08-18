import io
import mem
import os

main() {
	buffer := mem.alloc(6)
	fill(buffer)
	check(buffer)

	loop i := 0..buffer.len {
		writeByte(buffer.ptr[i])
	}

	mem.free(buffer)
}

fill(buffer string) {
	buffer.ptr[0] = 'H'
	buffer.ptr[1] = 'e'
	buffer.ptr[2] = 'l'
	buffer.ptr[3] = 'l'
	buffer.ptr[4] = 'o'
	buffer.ptr[5] = '\n'
}

check(buffer string) {
	assert buffer.ptr[0] == 'H'
	assert buffer.ptr[1] == 'e'
	assert buffer.ptr[2] == 'l'
	assert buffer.ptr[3] == 'l'
	assert buffer.ptr[4] == 'o'
	assert buffer.ptr[5] == '\n'
}

writeByte(n byte) {
	switch {
		n == 'H'  { io.write("H") }
		n == 'e'  { io.write("e") }
		n == 'l'  { io.write("l") }
		n == 'o'  { io.write("o") }
		n == '\n' { io.write("\n") }
		_         { os.exit(1) }
	}
}