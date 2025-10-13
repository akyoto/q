write(buffer string) -> (written int) {
	return syscall(_write, 1, buffer.ptr, buffer.len)
}

writeTo(fd int, buffer string) -> (written int, err error) {
	n := syscall(_write, fd, buffer.ptr, buffer.len)

	if n < 0 {
		return 0, n
	}

	return n, 0
}