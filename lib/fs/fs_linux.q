open(path *byte, flags int, mode int) -> !int {
	return syscall(_openat, -100, path, flags, mode)
}

size(fd int) -> int {
	stats := new(stat)
	syscall(_fstat, fd, stats)
	return stats.st_size
}

close(fd !int) -> int {
	return syscall(_close, fd)
}