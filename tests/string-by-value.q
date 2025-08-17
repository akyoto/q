import io
import mem

main() {
	len := 6
	ptr := mem.alloc(len)
	ptr[0] = 'H'
	ptr[1] = 'e'
	ptr[2] = 'l'
	ptr[3] = 'l'
	ptr[4] = 'o'
	ptr[5] = '\n'

	s := string{ptr: ptr, len: len}
	s.ptr[0] = 'W'
	s.ptr[1] = 'o'
	s.ptr[2] = 'r'
	s.ptr[3] = 'l'
	s.ptr[4] = 'd'

	io.write(s)
}