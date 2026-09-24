MyStruct {
	data [4]byte
}

main() {
	struct := new(MyStruct)
	struct.data[0] = 1
	struct.data[1] = 2
	struct.data[2] = 3
	struct.data[3] = 4
	assert struct.data[0] == 1
	assert struct.data[1] == 2
	assert struct.data[2] == 3
	assert struct.data[3] == 4
}