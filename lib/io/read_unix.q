read(buffer string) -> (count uint, err error) {
	return readFrom(stdin, buffer)
}

readFrom(fd uint, buffer string) -> (count uint, err error) {
	n := syscall(_read, fd, buffer.ptr, buffer.len)

	if n < 0 {
		return 0, n
	}

	return n, 0
}