main() {
	a := 1
	b := 2
	result := 0

	if true {
		c := b
		result = (8 & (c & a)) | (a << 1)
	}

	assert result == 2
}