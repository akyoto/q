read(buffer string) -> (read int) {
	return syscall(_read, 0, buffer.ptr, buffer.len)
}

readFrom(fd int, buffer string) -> (read int, err error) {
	n := syscall(_read, fd, buffer.ptr, buffer.len)

	if n < 0 {
		return 0, n
	}

	return n, 0
}