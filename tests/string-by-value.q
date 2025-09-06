import io
import mem

main() {
	buffer := mem.alloc(6)
	buffer[0] = 'H'
	buffer[1] = 'e'
	buffer[2] = 'l'
	buffer[3] = 'l'
	buffer[4] = 'o'
	buffer[5] = '\n'

	s := string{ptr: buffer.ptr, len: buffer.len}
	io.write(s)
	mem.free(buffer)
}