alloc(length uint) -> (buffer !string) {
	if length == 0 {
		return string{ptr: 0 as *uint8, len: 0}
	}

	aligned := (length + 15) & -16

	if heap.current + aligned < heap.next {
		x := heap.current
		heap.current += aligned
		return string{ptr: x, len: length}
	}

	size := (aligned + (pageSize - 1)) & -pageSize
	x := rawAlloc(size)
	heap.last = x
	heap.current = x + aligned
	heap.next = x + size
	return string{ptr: x, len: length}
}