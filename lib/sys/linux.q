read(fd Int, buffer Pointer, length Int) Int {
	require fd >= 0
	require buffer != 0
	require length >= 0
	ensure _ > -4096

	return syscall(0, fd, buffer, length)
}

write(fd Int, buffer Pointer, length Int) Int {
	require fd >= 0
	require buffer != 0
	require length >= 0
	ensure _ > -4096

	return syscall(1, fd, buffer, length)
}

open(fileName Text, flags Int, mode Int) Int {
	ensure _ > -4096

	return syscall(2, fileName, flags, mode)
}

close(fd Int) Int {
	require fd >= 0
	ensure _ > -4096

	return syscall(3, fd)
}

mmap(address Int, length Int, protection Int, flags Int) Int {
	require length > 0
	ensure _ > -4096

	return syscall(9, address, length, protection, flags)
}

munmap(address Pointer, length Int) Int {
	require address != 0
	require length > 0
	ensure _ > -4096
	ensure _ <= 0

	return syscall(11, address, length)
}

clone(flags Int, stackPointer Pointer) Int {
	ensure _ > -4096

	return syscall(56, flags, stackPointer)
}

exit(code Byte) {
	require code >= 0
	require code <= 125

	syscall(60, code)
}

getcwd(buffer Pointer, length Int) Int {
	require buffer != 0
	require length >= 0
	ensure _ > -4096

	return syscall(79, buffer, length)
}

chdir(path Text) Int {
	require path != 0
	ensure _ > -4096

	return syscall(80, path)
}

rename(old Text, new Text) Int {
	ensure _ > -4096

	return syscall(82, old, new)
}

mkdir(path Text, mode Int) Int {
	ensure _ > -4096

	return syscall(83, path, mode)
}

rmdir(path Text) Int {
	ensure _ > -4096

	return syscall(84, path)
}

unlink(fileName Text) Int {
	ensure _ > -4096

	return syscall(87, fileName)
}
