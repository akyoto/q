import slices

main() {
	a := new(int, 8)
	a[0] = 3
	a[1] = 2
	a[2] = 4
	a[3] = 1
	a[4] = 8
	a[5] = 7
	a[6] = 5
	a[7] = 6

	slices.sort(a)

	assert a[0] == 1
	assert a[1] == 2
	assert a[2] == 3
	assert a[3] == 4
	assert a[4] == 5
	assert a[5] == 6
	assert a[6] == 7
	assert a[7] == 8
}