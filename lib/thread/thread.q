import mem
import sys

create() {
	let stackSize = 8192
	let stack = mem.allocate(stackSize)
	let start = stack + stackSize - 8
	# *start = threadFunc
	sys.clone(stack, 2147585792)
}
