import mem
import sys

main() {
	# Allocate a few bytes
	length = 256
	buffer = mem.allocate(length)

	# Release the memory
	err = mem.free(buffer, length)
	sys.exit(err)
}
