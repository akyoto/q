import io

accept(fd uint) -> (conn !uint, err error) {
	conn := syscall(_accept, fd, 0, 0)

	if conn < 0 {
		return 0, conn
	}

	return conn, 0
}

close(fd !uint) -> error {
	return syscall(_close, fd)
}

connect(fd uint, address *any, length uint) -> error {
	return syscall(_connect, fd, address, length)
}

listen(fd uint, backlog int) -> error {
	return syscall(_listen, fd, backlog)
}

recv(fd uint, buffer string) -> (count uint, err error) {
	n, err := io.readFrom(fd, buffer)

	if err != 0 {
		return 0, err
	}

	return n, 0
}

socket(family int, type int, protocol int) -> (uint, error) {
	fd := syscall(_socket, family, type, protocol)

	if fd < 0 {
		return 0, fd
	}

	return fd, 0
}

send(fd uint, buffer string) -> (count uint, err error) {
	n, err := io.writeTo(fd, buffer)

	if err != 0 {
		return 0, err
	}

	return n, 0
}