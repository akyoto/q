read(fd, buffer, length) {
	ensure _ > -4096
	return syscall(0, fd, buffer, length)
}

write(fd, buffer, length) {
	ensure _ > -4096
	return syscall(1, fd, buffer, length)
}

open(fileName, flags, mode) {
	ensure _ > -4096
	return syscall(2, fileName, flags, mode)
}

close(fd) {
	ensure _ > -4096
	return syscall(3, fd)
}

exit(code) {
	require code >= 0
	require code <= 125
	syscall(60, code)
}

rename(old, new) {
	ensure _ > -4096
	return syscall(82, old, new)
}

mkdir(path, mode) {
	ensure _ > -4096
	return syscall(83, path, mode)
}

rmdir(path) {
	ensure _ > -4096
	return syscall(84, path)
}

unlink(fileName) {
	ensure _ > -4096
	return syscall(87, fileName)
}
