open(path *byte, flags int, mode int) -> int {
	return syscall(_open, path, flags, mode)
}

close(fd int) -> int {
	return syscall(_close, fd)
}