trimRight(s string) -> string {
	if s.len == 0 {
		return s
	}

	end := s.len

	loop {
		char := s[end-1]

		if char != ' ' && char != '\t' && char != '\n' && char != '\r' {
			return s[..end]
		}

		end -= 1

		if end == 0 {
			return ""
		}
	}
}