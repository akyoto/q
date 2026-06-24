import run

free(buffer !string) {
	if buffer.ptr == 0 {
		return
	}

	if buffer.ptr + buffer.len == heap.current {
		zero(buffer)
		heap.current -= buffer.len

		if heap.current < heap.last {
			size := (heap.next - heap.last) as uint
			osFree(heap.last, size)
			heap.next = heap.last
			heap.last -= pageSize
		}

		return
	}

	run.crash()
}