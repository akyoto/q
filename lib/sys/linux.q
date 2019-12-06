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

mmap(address, length, protection, flags) {
	ensure _ > -4096
	return syscall(9, address, length, protection, flags)
}

# mmap(address, length, protection, flags, fd, offset) {
# 	ensure _ > -4096
# 	return syscall(9, address, length, protection, flags, fd, offset)
# }

munmap(address, length) {
	ensure _ > -4096
	ensure _ <= 0
	return syscall(11, address, length)
}

clone(flags, stackPointer) {
	ensure _ > -4096
	return syscall(56, flags, stackPointer)
}

# clone(flags, stackPointer, parentThreadID, childThreadID) {
# 	ensure _ > -4096
# 	return syscall(56, flags, stackPointer, parentThreadID, childThreadID)
# }

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
