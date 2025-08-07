main() {
	i := 10

	loop 0..i {
		i -= 1
	}

	assert i == 5
}