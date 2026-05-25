copy(to string, from string) {
	len := from.len

	if to.len < len {
		len = to.len
	}

	copy(to.ptr, from.ptr, len)
}

copy(to *byte, from *byte, len uint) {
	loop i := 0..len {
		to[i] = from[i]
	}
}