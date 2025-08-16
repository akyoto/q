read2(buffer *byte, length int) -> int {
	return syscall(_read, 0, buffer, length)
}

write2(buffer *byte, length int) -> (written int) {
	return syscall(_write, 1, buffer, length)
}