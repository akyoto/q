main() {
	x := new(int32)
	[x] = -1
	assert [x] == -1
	[x] = -2147483647
	assert [x] == -2147483647
	[x] = -2147483648
	assert [x] == -2147483648
}