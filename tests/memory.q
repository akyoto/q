import mem

main() {
	length := 6
	address := mem.alloc(length)
	fill(address)
	check(address)
	mem.free(address, length)
}

fill(address *byte) {
	address[0] = 'H'
	address[1] = 'e'
	address[2] = 'l'
	address[3] = 'l'
	address[4] = 'o'
	address[5] = '\n'
}

check(address *byte) {
	assert address[0] == 'H'
	assert address[1] == 'e'
	assert address[2] == 'l'
	assert address[3] == 'l'
	assert address[4] == 'o'
	assert address[5] == '\n'
}