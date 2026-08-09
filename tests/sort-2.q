import slices

main() {
	a := new(int, 2)
	a[0] = 2
	a[1] = 1

	slices.sort(a)

	assert a[0] == 1
	assert a[1] == 2
}

