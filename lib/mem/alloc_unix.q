import run

alloc(length int) -> (address *byte) {
	x := mmap(0, length, read|write, private|anonymous)

	if x < 0x1000 {
		run.crash()
	}

	return x
}

free(address *any, length int) {
	munmap(address, length)
}