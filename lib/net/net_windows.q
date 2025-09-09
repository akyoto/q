accept(fd int) -> (int, error) {
	conn := ws2_32.accept(fd, 0, 0)

	if conn < 0 {
		return 0, conn
	}

	return conn, 0
}

close(fd int) -> error {
	ws2_32.shutdown(fd, 1)
	ws2_32.recv(fd, 0, 0, 0)
	return ws2_32.closesocket(fd)
}

listen(fd int, backlog int) -> error {
	return ws2_32.listen(fd, backlog)
}

recv(fd int, buffer string) -> int {
	return ws2_32.recv(fd, buffer.ptr, buffer.len as int, 0)
}

socket(family int, type int, protocol int) -> (int, error) {
	s := ws2_32.socket(family, type, protocol)

	if s < 0 {
		return 0, s
	}

	return s, 0
}

send(fd int, buffer string) -> int {
	return ws2_32.send(fd, buffer.ptr, buffer.len as int, 0)
}

extern {
	ws2_32 {
		accept(fd int, address *any|nil, length int) -> int
		bind(socket int, address *any, length int) -> (error int)
		closesocket(fd int) -> int
		listen(fd int, backlog int) -> int
		recv(fd int, address *byte|nil, length int, flags int) -> int
		socket(family int, type int, protocol int) -> int
		send(fd int, address *byte, length int, flags int) -> int
		shutdown(fd int, how int) -> int
	}
}