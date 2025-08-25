open(path *byte, flags int, mode int) -> !int {
	return syscall(_open, path, flags, mode)
}

size(fd int) -> int {
	stats := new(stat)
	syscall(_fstat64, fd, stats)
	return stats.st_size
}

close(fd !int) -> int {
	return syscall(_close, fd)
}