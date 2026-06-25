openRead(path *byte) -> (!uint, error) {
	fd := syscall(_openat, -100, path, readOnly, 0)

	if fd < 0 {
		return 0, fd
	}

	return fd, 0
}

openWrite(path *byte) -> (!uint, error) {
	fd := syscall(_openat, -100, path, writeOnly | create | truncate, 0o644)

	if fd < 0 {
		return 0, fd
	}

	return fd, 0
}

size(fd uint) -> (uint, error) {
	stats := new(FileStat)
	err := syscall(_fstat, fd, stats)

	if err != 0 {
		return 0, err
	}

	size := stats.size as uint
	return size, 0
}

close(fd !uint) -> error {
	return syscall(_close, fd)
}

memfd_create(path *byte, flags uint) -> (!uint, error) {
	fd := syscall(_memfd_create, path, flags)

	if fd < 0 {
		return 0, fd
	}

	return fd, 0
}

ftruncate(fd uint, length uint) -> error {
	return syscall(_ftruncate, fd, length)
}