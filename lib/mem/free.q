free(buffer ![]byte) {
	if buffer.len == 0 {
		return
	}

	if heap.current == heap.min {
		freePtr := heap.min - 32
		freeSize := (heap.max - freePtr) as uint
		heap = [freePtr as *Heap]
		rawFree(freePtr, freeSize)
	}

	aligned := (buffer.len + 15) & -16
	assert buffer.ptr + aligned == heap.current
	zero(buffer)
	heap.current -= aligned
}