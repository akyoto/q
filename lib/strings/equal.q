equal(a string, b string) -> bool {
	if a.len != b.len {
		return false
	}

	if a.ptr == b.ptr {
		return true
	}

	loop i := 0..a.len {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}