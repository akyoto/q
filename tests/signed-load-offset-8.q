main() {
	a := new(int8, 3)
	a[0] = -1
	a[1] = -127
	a[2] = -128
	assert a[0] == -1
	assert a[1] == -127
	assert a[2] == -128
	delete(a)
}