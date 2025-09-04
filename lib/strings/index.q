index(a string, b string) -> (uint, error) {
	if b.len > a.len {
		return 0, -1
	}

	if b.len == 0 {
		return 0, 0
	}

	j := 0

	loop i := 0..a.len {
		if a[i] == b[j] {
			j += 1

			if j == b.len {
				return i - j + 1, 0
			}
		} else {
			j = 0
		}
	}

	return 0, -1
}