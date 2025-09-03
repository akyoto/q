import io

accept(fd int) -> (int, error) {
	conn := syscall(_accept, fd, 0, 0)

	if conn < 0 {
		return 0, conn
	}

	return conn, 0
}

close(fd int) -> error {
	return syscall(_close, fd)
}

listen(fd int, backlog int) -> error {
	return syscall(_listen, fd, backlog)
}

recv(fd int, buffer string) -> int {
	return io.readFrom(fd, buffer)
}

socket(family int, type int, protocol int) -> (int, error) {
	s := syscall(_socket, family, type, protocol)

	if s < 0 {
		return 0, s
	}

	return s, 0
}

send(fd int, buffer string) -> int {
	return io.writeTo(fd, buffer)
}