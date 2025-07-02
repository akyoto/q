write(buffer string) -> (written int) {
	return syscall(64, 1, buffer.ptr, buffer.len)
}