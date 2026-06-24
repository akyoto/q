import run

osAlloc(length uint) -> *uint8 {
	if length == 0 {
		return 0 as *uint8
	}

	x := mmap(0, length, read|write, private|anonymous, -1, 0)

	if x < 0x1000 {
		run.crash()
	}

	return x
}

osFree(ptr *any, len uint) {
	if ptr == 0 {
		return
	}

	munmap(ptr, len)
}