zero(buffer string) {
	loop i := 0..buffer.len {
		buffer.ptr[i] = 0
	}
}