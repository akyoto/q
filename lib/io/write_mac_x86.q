write(buffer string) -> (written int) {
	return syscall(0x2000004, 1, buffer.ptr, buffer.len)
}