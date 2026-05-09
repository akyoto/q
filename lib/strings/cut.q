cut(s string, separator string) -> (before string, after string, err error) {
	pos, err := index(s, separator)

	if err != 0 {
		return "", "", -1
	}

	return s[..pos], s[pos+separator.len..], 0
}