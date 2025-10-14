openRead(path *byte) -> (!int, error) {
	fd := syscall(_open, path, readOnly, 0)

	if fd < 0 {
		return 0, fd
	}

	return fd, 0
}

openWrite(path *byte) -> (!int, error) {
	fd := syscall(_open, path, writeOnly | create | truncate, 0o644)

	if fd < 0 {
		return 0, fd
	}

	return fd, 0
}

size(fd int) -> (uint, error) {
	stats := new(FileStat)
	err := syscall(_fstat64, fd, stats)

	if err != 0 {
		delete(stats)
		return 0, err
	}

	size := stats.size as uint
	delete(stats)
	return size, 0
}

close(fd !int) -> error {
	return syscall(_close, fd)
}