main() {
	a := new(int8, 3)
	a[0] = -1
	a[1] = -2
	a[2] = -3
	a[3] = -4
	assert a[0] == -1
	assert a[1] == -2
	assert a[2] == -3
	assert a[3] == -4
	delete(a)
}