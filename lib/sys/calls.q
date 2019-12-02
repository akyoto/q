read(fd, buffer, length) {
	return syscall(0, fd, buffer, length)
}

write(fd, buffer, length) {
	return syscall(1, fd, buffer, length)
}

open(fileName, flags, mode) {
	return syscall(2, fileName, flags, mode)
}

close(fd) {
	return syscall(3, fd)
}

exit(code) {
	syscall(60, code)
}

rename(old, new) {
	return syscall(82, old, new)
}

mkdir(path, mode) {
	return syscall(83, path, mode)
}

rmdir(path) {
	return syscall(84, path)
}

unlink(fileName) {
	return syscall(87, fileName)
}
