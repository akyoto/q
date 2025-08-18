open(path *byte, flags int, mode int) -> int {
	return syscall(_openat, -100, path, flags, mode)
}

close(fd int) -> int {
	return syscall(_close, fd)
}