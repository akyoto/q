read(fd, buffer, length) {
	require fd >= 0
	require buffer != 0
	require length >= 0
	ensure _ > -4096

	return syscall(0, fd, buffer, length)
}

write(fd, buffer, length) {
	require fd >= 0
	require buffer != 0
	require length >= 0
	ensure _ > -4096

	return syscall(1, fd, buffer, length)
}

open(fileName, flags, mode) {
	ensure _ > -4096

	return syscall(2, fileName, flags, mode)
}

close(fd) {
	require fd >= 0
	ensure _ > -4096

	return syscall(3, fd)
}

mmap(address, length, protection, flags) {
	require length > 0
	ensure _ > -4096

	return syscall(9, address, length, protection, flags)
}

munmap(address, length) {
	require address != 0
	require length > 0
	ensure _ > -4096
	ensure _ <= 0

	return syscall(11, address, length)
}

clone(flags, stackPointer) {
	ensure _ > -4096

	return syscall(56, flags, stackPointer)
}

exit(code) {
	require code >= 0
	require code <= 125

	syscall(60, code)
}

getcwd(buffer, length) {
	require buffer != 0
	require length >= 0
	ensure _ > -4096

	return syscall(79, buffer, length)
}

chdir(path) {
	require path != 0
	ensure _ > -4096

	return syscall(80, path)
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
