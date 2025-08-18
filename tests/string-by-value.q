import io
import mem

main() {
	buffer := mem.alloc(6)
	buffer.ptr[0] = 'H'
	buffer.ptr[1] = 'e'
	buffer.ptr[2] = 'l'
	buffer.ptr[3] = 'l'
	buffer.ptr[4] = 'o'
	buffer.ptr[5] = '\n'

	s := string{ptr: buffer.ptr, len: buffer.len}
	s.ptr[0] = 'W'
	s.ptr[1] = 'o'
	s.ptr[2] = 'r'
	s.ptr[3] = 'l'
	s.ptr[4] = 'd'

	io.write(s)
}