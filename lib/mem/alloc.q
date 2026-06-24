alloc(length uint) -> (buffer !string) {
	if length == 0 {
		return string{ptr: 0 as *uint8, len: 0}
	}

	if heap.current + length < heap.next {
		ptr := heap.current
		heap.current += length
		return string{ptr: ptr, len: length}
	}

	size := (length + (pageSize - 1)) & -pageSize
	x := rawAlloc(size)
	heap.last = x
	heap.current = x + length
	heap.next = x + size
	return string{ptr: x, len: length}
}