write(fd int, buffer *byte, length int) -> int {
	return syscall(1, fd, buffer, length)
}