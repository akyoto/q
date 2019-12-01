open(fileName) {
	return syscall(2, fileName, 66, 438)
}

close(fd) {
	return syscall(3, fd)
}

write(fd, msg, length) {
	return syscall(1, fd, msg, length)
}

unlink(fileName) {
	return syscall(87, fileName)
}
