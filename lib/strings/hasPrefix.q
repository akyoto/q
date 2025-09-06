hasPrefix(s string, prefix string) -> bool {
	if s.len < prefix.len {
		return false
	}

	loop i := 0..prefix.len {
		if s[i] != prefix[i] {
			return false
		}
	}

	return true
}