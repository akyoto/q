main() {
	i := 10

	for 0..i {
		i -= 1
	}

	assert i == 5
}