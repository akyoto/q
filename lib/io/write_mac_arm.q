write(buffer string) -> (written int) {
	return syscall(0x4, 1, buffer.ptr, buffer.len)
}