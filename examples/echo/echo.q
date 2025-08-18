import io
import mem

main() {
	buffer := mem.alloc(4096)

	loop {
		n := io.read(buffer)

		if n <= 0 {
			mem.free(buffer)
			return
		}

		io.write(string{ptr: buffer.ptr, len: n})
	}
}