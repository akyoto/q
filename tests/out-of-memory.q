import mem

main() {
	address := mem.alloc(0x4000000000000)
	assert address == 0
}