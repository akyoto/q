main() {
	total := 0

	loop i := 0..4 {
		total += i
	}

	total += 1
	total += 2
	assert 1 == 1
	total += 3
	total += 4
}