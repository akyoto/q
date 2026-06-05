main() {
	x := new(int16)
	[x] = -1
	assert [x] == -1
	[x] = -32767
	assert [x] == -32767
	[x] = -32768
	assert [x] == -32768
	delete(x)
}