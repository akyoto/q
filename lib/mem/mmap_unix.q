mmap(address int, length uint, protection int, flags int, fd int, offset int) -> (address *byte) {
	return syscall(_mmap, address, length, protection, flags, fd, offset)
}

munmap(address *any, length uint) -> int {
	return syscall(_munmap, address, length)
}