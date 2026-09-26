global {
	data [10]int
}

main() {
	data[1] = 11
	a := data[1..4]
	assert a[0] == 11
}