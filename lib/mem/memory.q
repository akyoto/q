import sys

allocate(length Int) -> Pointer {
	return sys.mmap(0, length, 3, 290)
}

free(pointer Pointer, length Int) -> Int {
	return sys.munmap(pointer, length)
}
