main() {
	a := new(uint32)
	[a] = 1

	assert cas(a, 1, 2)
	assert [a] == 2

	assert !cas(a, 1, 3)
	assert [a] == 2

	assert cas(a, 2, 1)
	assert [a] == 1
}