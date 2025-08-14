import mem

main() {
	length := 6
	address := mem.alloc(length)
	assert address != 0
	fill(address)
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