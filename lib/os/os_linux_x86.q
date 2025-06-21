write(fd int, buffer *byte, length int) -> (written int) {
	return syscall(1, fd, buffer, length)
}