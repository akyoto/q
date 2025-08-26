import mem

main() {
	a := mem.alloc(4)

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

	mem.free(a)
}