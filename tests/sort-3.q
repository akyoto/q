import slices

main() {
	a := new(int, 3)
	a[0] = 2
	a[1] = 3
	a[2] = 1

	slices.sort(a)

	assert a[0] == 1
	assert a[1] == 2
	assert a[2] == 3
}

