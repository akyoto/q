import mem
import sys

main() {
	# Allocate a few bytes
	let length = 256
	let buffer = mem.allocate(length)

	# Free the memory
	let err = mem.free(buffer, length)
	sys.exit(err)
}
