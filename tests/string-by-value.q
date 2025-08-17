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
	io.write(s)
}