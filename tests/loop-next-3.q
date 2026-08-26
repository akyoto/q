main() {
	assert f() == 7
}

f() -> int {
	count := 0

	loop i := 0..20 {
		count += 1

		if i == 2 {
			loop.next()
		}

		if count == 7 {
			loop.stop()
		}
	}

	return count
}