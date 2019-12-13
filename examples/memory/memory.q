import mem
import sys

main() {
	# Allocate a few bytes
	let length = 256
	let buffer = mem.allocate(length)

	store(buffer, 0, 1, 65)
	store(buffer, 1, 1, 66)
	store(buffer, 2, 1, 67)
	store(buffer, 3, 1, 68)
	store(buffer, 4, 1, 10)
	sys.write(1, buffer, 5)

	# Free the memory
	let err = mem.free(buffer, length)
	sys.exit(err)
}
