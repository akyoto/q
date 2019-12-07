import mem
import sys

create() {
	stackSize = 8192
	stack = mem.allocate(stackSize)
	start = stack + stackSize - 8
	# *start = threadFunc
	sys.clone(stack, 2147585792)
}

threadFunc() {

}
