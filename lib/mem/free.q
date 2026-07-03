import run

free(buffer ![]byte) {
	if buffer.ptr == 0 {
		return
	}

	aligned := (buffer.len + 15) & -16

	if buffer.ptr + aligned != heap.current {
		run.crash()
	}

	zero(buffer)
	heap.current -= aligned
}