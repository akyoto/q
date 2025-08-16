import io
import mem

main() {
	length := 4096
	address := mem.alloc(length)

	loop {
		n := io.read2(address, length)

		if n <= 0 {
			mem.free(address, length)
			return
		}

		io.write2(address, n)
	}
}