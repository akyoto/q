alloc(length uint) -> (buffer ![]byte) {
	if length == 0 {
		return []byte{ptr: 0 as *byte, len: 0}
	}

	aligned := (length + 15) & -16

	if heap.current + aligned > heap.max {
		size := (heap.max - heap.min + 32) * 2 as uint

		if size < pageSize {
			size = (aligned + (pageSize - 1)) & -pageSize
		}

		x := rawAlloc(size)
		[x as *Heap] = heap
		current := x + 32
		heap.min = current
		heap.current = current + aligned
		heap.max = x + size
		return []byte{ptr: current, len: length}
	}

	x := heap.current
	heap.current += aligned
	return []byte{ptr: x, len: length}
}