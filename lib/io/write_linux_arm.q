write(buffer []byte) -> (written int) {
	return syscall(64, 1, buffer, len(buffer))
}