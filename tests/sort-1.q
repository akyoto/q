import slices

main() {
	a := new(int, 1)
	a[0] = 1

	slices.sort(a)

	assert a[0] == 1
}