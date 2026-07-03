const {
	pageSize = 0x200000
}

global {
	heap Heap
}

Heap {
	min *byte
	current *byte
	max *byte
}