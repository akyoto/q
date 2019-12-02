write(fd, msg, length) {
	return syscall(1, fd, msg, length)
}

open(fileName, flags, mode) {
	return syscall(2, fileName, flags, mode)
}

close(fd) {
	return syscall(3, fd)
}

unlink(fileName) {
	return syscall(87, fileName)
}
