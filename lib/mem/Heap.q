const {
	pageSize = 0x200000
}

global {
	heap Heap
}

Heap {
	last *uint8
	current *uint8
	next *uint8
}