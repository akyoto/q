write(buffer string) -> (written int) {
	return syscall(1, 1, buffer.ptr, buffer.len)
}