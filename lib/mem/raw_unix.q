import run

rawAlloc(length uint) -> *uint8 {
	x := mmap(0, length, read|write, private|anonymous, -1, 0)

	if (x as int) < 0x1000 {
		run.crash()
	}

	return x
}

rawFree(ptr *any, len uint) {
	munmap(ptr, len)
}