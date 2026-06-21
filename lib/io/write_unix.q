write(buffer string) -> (count uint, err error) {
	return writeTo(stdout, buffer)
}

writeTo(fd int, buffer string) -> (count uint, err error) {
	n := syscall(_write, fd, buffer.ptr, buffer.len)

	if n < 0 {
		return 0, n
	}

	return n, 0
}