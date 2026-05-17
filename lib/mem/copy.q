copy(to string, from string) {
	len := from.len

	if to.len < len {
		len = to.len
	}

	loop i := 0..len {
		to[i] = from[i]
	}
}