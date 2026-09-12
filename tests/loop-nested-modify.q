main() {
	count := 0

	loop i := 2..11 {
		loop j := 2..3 {
			j += 2
			i += 2
		}

		count += 1
	}

	assert count == 3
}