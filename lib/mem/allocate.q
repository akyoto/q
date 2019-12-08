import sys

allocate(length) -> Pointer {
	return sys.mmap(0, length, 3, 290)
}
