accept(fd int, address *any, length int) -> int {
	return syscall(_accept, fd, address, length)
}

close(fd int) -> int {
	return syscall(_close, fd)
}

listen(fd int, backlog int) -> int {
	return syscall(_listen, fd, backlog)
}

socket(family int, type int, protocol int) -> int {
	return syscall(_socket, family, type, protocol)
}