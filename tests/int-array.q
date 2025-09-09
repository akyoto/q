import mem

main() {
	n := 10
	buf := mem.alloc(n * 8)
	buf2 := IntArray{ptr: buf.ptr as *int, len: n}

	loop i := 0..buf2.len {
		buf2[i] = i * 1234567890
	}

	loop i := 0..buf2.len {
		assert buf2[i] == i * 1234567890
	}

	mem.free(buf)
}

IntArray {
	ptr *int
	len uint
}