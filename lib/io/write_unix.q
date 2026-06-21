write(buffer string) -> (count int, err error) {
	return writeTo(1, buffer)
}

writeTo(fd int, buffer string) -> (count int, err error) {
	n := syscall(_write, fd, buffer.ptr, buffer.len)

	if n < 0 {
		return 0, n
	}

	return n, 0
}