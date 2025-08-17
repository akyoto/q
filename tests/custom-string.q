import io
import mem

String {
	ptr *byte
	len int
}

main() {
	s := new(String)
	s.len = 6
	s.ptr = mem.alloc(s.len)

	loop i := 0..s.len {
		s.ptr[i] = '0' + i
	}

	print(s)
}

print(s *String) {
	io.write2(s.ptr, s.len)
}