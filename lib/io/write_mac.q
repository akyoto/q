write(buffer []byte) -> (written int) {
	return syscall(0x2000004, 1, buffer, len(buffer))
}