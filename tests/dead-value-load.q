main() {
	v := 4
	a := new(int, 4)
	a[0] = -39
	a[1] = 28
	a[2] = 22
	a[3] = -31
	sum := 0
	sum += a[0]
	sum += a[1]
	sum += a[2]
	sum += a[3]
	sel := 0

	switch {
		v > 2 { sel = 7 }
		_     { sel = -3 }
	}

	assert sel == 7
}