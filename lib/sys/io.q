read(fd, buffer, length) {
	return syscall(0, fd, buffer, length)
}

write(fd, buffer, length) {
	return syscall(1, fd, buffer, length)
}
