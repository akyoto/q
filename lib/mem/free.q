import sys

free(pointer, length) {
	return sys.munmap(pointer, length)
}
