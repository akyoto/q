import run

alloc(length uint) -> (buffer !string) {
	x := mmap(0, length, read|write, private|anonymous, -1, 0)

	if x < 0x1000 {
		run.crash()
	}

	return string{ptr: x, len: length}
}

free(buffer !string) {
	munmap(buffer.ptr, buffer.len)
}