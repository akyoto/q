MyStruct {
	a [4]byte
	b [4]byte
	c [4]byte
}

main() {
	struct := new(MyStruct)

	struct.a[0] = 1
	struct.a[1] = 2
	struct.a[2] = 3
	struct.a[3] = 4

	struct.b[0] = 5
	struct.b[1] = 6
	struct.b[2] = 7
	struct.b[3] = 8

	struct.c[0] = 9
	struct.c[1] = 10
	struct.c[2] = 11
	struct.c[3] = 12

	assert struct.a[0] == 1
	assert struct.a[1] == 2
	assert struct.a[2] == 3
	assert struct.a[3] == 4

	assert struct.b[0] == 5
	assert struct.b[1] == 6
	assert struct.b[2] == 7
	assert struct.b[3] == 8

	assert struct.c[0] == 9
	assert struct.c[1] == 10
	assert struct.c[2] == 11
	assert struct.c[3] == 12
}