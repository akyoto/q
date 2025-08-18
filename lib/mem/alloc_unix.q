import run

alloc(length int) -> (buffer string) {
	x := mmap(0, length, read|write, private|anonymous)

	if x < 0x1000 {
		run.crash()
	}

	return string{ptr: x, len: length}
}

free(buffer string) {
	munmap(buffer.ptr, buffer.len)
}