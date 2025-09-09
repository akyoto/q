open(path *byte, flags int, mode int) -> (!int, error) {
	fd := syscall(_openat, -100, path, flags, mode)

	if fd < 0 {
		return 0, fd
	}

	return fd, 0
}

size(fd int) -> (uint, error) {
	stats := new(stat)
	err := syscall(_fstat, fd, stats)

	if err != 0 {
		delete(stats)
		return 0, err
	}

	size := stats.st_size as uint
	delete(stats)
	return size, 0
}

close(fd !int) -> error {
	return syscall(_close, fd)
}