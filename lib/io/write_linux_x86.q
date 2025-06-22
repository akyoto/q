write(buffer []byte) -> (written int) {
	return syscall(1, 0, buffer, len(buffer))
}