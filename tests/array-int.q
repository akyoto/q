main() {
	x := new(int, 10)
	assert x.len == 10

	loop i := 0..x.len {
		x[i] = i * 1234567890
	}

	loop i := 0..x.len {
		assert x[i] == i * 1234567890
	}
}