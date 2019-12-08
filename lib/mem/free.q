import sys

free(pointer Pointer, length Int) -> Int {
	return sys.munmap(pointer, length)
}
