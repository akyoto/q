main() {
	x := new(int8)
	[x] = -1
	assert [x] == -1
	[x] = -127
	assert [x] == -127
	[x] = -128
	assert [x] == -128
	delete(x)
}