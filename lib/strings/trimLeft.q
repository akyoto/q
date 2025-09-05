trimLeft(s string) -> string {
	if s.len == 0 {
		return s
	}

	start := 0

	loop {
		char := s[start]

		if char != ' ' && char != '\t' && char != '\n' && char != '\r' {
			return s[start..]
		}

		start += 1

		if start == s.len {
			return ""
		}
	}
}