import sys

allocate(length) {
	return sys.mmap(0, length, 3, 290)
}

free(pointer, length) {
	return sys.munmap(pointer, length)
}
