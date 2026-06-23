main() {
	a := new(int16, 3)
	a[0] = -1
	a[1] = -32767
	a[2] = -32768
	assert a[0] == -1
	assert a[1] == -32767
	assert a[2] == -32768
}