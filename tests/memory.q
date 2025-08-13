import mem

main() {
	length := 1024
	address := mem.alloc(length)
	assert address != 0
	mem.free(address, length)
}