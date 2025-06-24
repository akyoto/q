write(buffer []byte) -> (written int) {
	return syscall(1, 1, buffer, len(buffer))
}