length(ptr *byte) -> uint {
	len := 0

	loop {
		if ptr[len] == 0 {
			return len
		}

		len += 1
	}
}