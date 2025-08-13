mmap(address int, length uint, protection int, flags int) -> *any {
	return syscall(_mmap, address, length, protection, flags, 0, 0)
}

munmap(address *any, length uint) -> int {
	return syscall(_munmap, address, length)
}

const {
	read 0x1
	write 0x2
	private 0x02
	anonymous 0x1000
}