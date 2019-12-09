import sys

allocate(length Int) -> Pointer {
	return sys.mmap(0, length, 3, 290)
}
