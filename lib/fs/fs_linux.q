open(path *byte, flags int, mode int) -> (!int, error) {
	fd := syscall(_openat, -100, path, flags, mode)

	if fd < 0 {
		return 0, fd
	}

	return fd, 0
}

size(fd int) -> (int, error) {
	stats := new(stat)
	err := syscall(_fstat, fd, stats)

	if err != 0 {
		return 0, err
	}

	return stats.st_size, 0
}

close(fd !int) -> error {
	return syscall(_close, fd)
}