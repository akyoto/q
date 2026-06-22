import run

alloc(length uint) -> (buffer !string) {
	if length == 0 {
		return string{ptr: 0 as *uint8, len: 0}
	}

	x := mmap(0, length, read|write, private|anonymous, -1, 0)

	if x < 0x1000 {
		run.crash()
	}

	return string{ptr: x, len: length}
}

free(buffer !string) {
	if buffer.ptr == 0 {
		return
	}

	munmap(buffer.ptr, buffer.len)
}