write16(buffer string, v uint16) {
	[buffer.ptr as *uint16] = v
}

write32(buffer string, v uint32) {
	[buffer.ptr as *uint32] = v
}