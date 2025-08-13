alloc(length int) -> (address *any) {
	x := mmap(0, length, read|write, private|anonymous)

	if x < 0x1000 {
		return 0
	}

	return x
}

free(address *any, length int) {
	munmap(address, length)
}