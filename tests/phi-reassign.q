main() {
	a := 0
	b := 1
	c := 2
	d := 3
	e := 4
	f := 5

	if true {
		a += 1
		b += 1
		c += 1
		d += 1
		e += 1
		f += 1
	}

	assert a == 1
	assert b == 2
	assert c == 3
	assert d == 4
	assert e == 5
	assert f == 6
}