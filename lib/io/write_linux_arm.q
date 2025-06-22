write(buffer []byte) -> (written int) {
	return syscall(64, 0, buffer, len(buffer))
}