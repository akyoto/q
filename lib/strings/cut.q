cut(s string, separator string) -> (string, string) {
	pos, err := index(s, separator)

	if err != 0 {
		return s, ""
	}

	return s[..pos], s[pos+separator.len..]
}