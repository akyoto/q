main() {
	i := 0

	loop 0..10 {
		i += 1
	}

	assert i == 10
}