hasSuffix(s string, suffix string) -> bool {
	if s.len < suffix.len {
		return false
	}

	start := s.len - suffix.len

	loop i := start..s.len {
		if s[i] != suffix[i-start] {
			return false
		}
	}

	return true
}