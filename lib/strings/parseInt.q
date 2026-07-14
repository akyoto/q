import errors

parseInt(s string) -> (int, error) {
	if s.len == 0 {
		return 0, errors.invalidArgument
	}

	scalar := 1

	if s[0] == '-' {
		scalar = -1
		s = s[1..]
	}

	ptr := s.ptr
	end := s.ptr + s.len
	result := 0

	loop {
		if ptr >= end {
			return result * scalar, 0
		}

		c := [ptr]

		if c >= '0' && c <= '9' {
			result = result * 10 + (c - '0')
			ptr += 1
			loop.next()
		}

		return 0, errors.invalidArgument
	}
}