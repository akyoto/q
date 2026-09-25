main() {
	a0 := 0
	a1 := 0
	a2 := 0
	a3 := 0
	a4 := 0
	a5 := 0
	a6 := 0

	loop i := 0..2 {
		a0 += 1
		a1 += 1
		a2 += 1
		a3 += 1
		a4 += 1
		a5 += 1
		a6 += 1
	}

	assert a0 == 2
	assert a1 == 2
	assert a2 == 2
	assert a3 == 2
	assert a4 == 2
	assert a5 == 2
	assert a6 == 2
}
