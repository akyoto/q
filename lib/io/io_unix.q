write(buffer string) -> (written int) {
	return syscall(_write, 1, buffer.ptr, buffer.len)
}