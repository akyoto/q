main() {
	assert f() == 4
}

f() -> int {
	count := 0

	loop i := 0..4 {
		count += 1

		if i == 2 {
			loop.next()
		}
	}

	return count
}