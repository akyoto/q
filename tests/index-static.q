import mem

main() {
	buffer := mem.alloc(4)
	a := buffer.ptr

	assert a[0] == 0
	assert a[1] == 0
	assert a[2] == 0
	assert a[3] == 0

	a[0] = 0
	a[1] = 1
	a[2] = 2
	a[3] = 3

	assert a[0] == 0
	assert a[1] == 1
	assert a[2] == 2
	assert a[3] == 3
}