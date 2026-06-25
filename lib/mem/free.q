import run

free(buffer !string) {
	if buffer.ptr == 0 {
		return
	}

	aligned := (buffer.len + 15) & -16

	if buffer.ptr + aligned == heap.current {
		zero(buffer)
		heap.current -= aligned

		if heap.current < heap.last {
			size := (heap.next - heap.last) as uint
			rawFree(heap.last, size)
			heap.next = heap.last
			heap.last -= pageSize
		}

		return
	}

	run.crash()
}