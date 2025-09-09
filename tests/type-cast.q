import mem

main() {
	buffer := mem.alloc(4)
	intPtr := buffer.ptr as *int
	intPtr[0] = 0x12345678
	assert buffer[0] == 0x78
	assert buffer[1] == 0x56
	assert buffer[2] == 0x34
	assert buffer[3] == 0x12
	mem.free(buffer)
}