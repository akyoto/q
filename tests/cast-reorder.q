main() {
	a := -1952
	assert ((a % 11 + (a as uint32) >> 27) as int) == -1
}