global {
	data [4]byte
}

main() {
	data[0] = 1
	data[1] = 2
	data[2] = 3
	data[3] = 4

	assert data[0] == 1
	assert data[1] == 2
	assert data[2] == 3
	assert data[3] == 4
}