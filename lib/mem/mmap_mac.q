mmap(address int, length uint, protection int, flags int) -> (address *byte) {
	return syscall(_mmap, address, length, protection, flags, 0, 0)
}

munmap(address *any, length uint) -> int {
	return syscall(_munmap, address, length)
}