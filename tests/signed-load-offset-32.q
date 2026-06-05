main() {
	a := new(int32, 3)
	a[0] = -1
	a[1] = -2147483647
	a[2] = -2147483648
	assert a[0] == -1
	assert a[1] == -2147483647
	assert a[2] == -2147483648
	delete(a)
}