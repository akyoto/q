main() {
	a := new(uint32)
	[a] = 1

	b := cas(a, 1, 2)
	assert b == 1
	assert [a] == 2

	b = cas(a, 1, 3)
	assert b == 2
	assert [a] == 2

	b = cas(a, 2, 1)
	assert b == 2
	assert [a] == 1
}