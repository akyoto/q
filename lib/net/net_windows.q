accept(fd uint) -> (conn !uint, err error) {
	conn := ws2_32.accept(fd, 0, 0)

	if conn == SOCKET_ERROR {
		return 0, ws2_32.WSAGetLastError()
	}

	return conn, 0
}

close(fd !uint) -> error {
	if ws2_32.shutdown(fd, 1) == SOCKET_ERROR {
		return ws2_32.WSAGetLastError()
	}

	if ws2_32.recv(fd, 0, 0, 0) == SOCKET_ERROR {
		return ws2_32.WSAGetLastError()
	}

	if ws2_32.closesocket(fd) == SOCKET_ERROR {
		return ws2_32.WSAGetLastError()
	}

	return 0
}

connect(fd uint, address *any, length uint) -> error {
	err := ws2_32.connect(fd, address, length as int)

	if err == SOCKET_ERROR {
		return ws2_32.WSAGetLastError()
	}

	return 0
}

listen(fd uint, backlog int) -> error {
	err := ws2_32.listen(fd, backlog)

	if err == SOCKET_ERROR {
		return ws2_32.WSAGetLastError()
	}

	return 0
}

recv(fd uint, buffer string) -> (count uint, err error) {
	n := ws2_32.recv(fd, buffer.ptr, buffer.len as int, 0)

	if n == SOCKET_ERROR {
		return 0, ws2_32.WSAGetLastError()
	}

	return n as uint, 0
}

socket(family int, type int, protocol int) -> (uint, error) {
	fd := ws2_32.socket(family, type, protocol)

	if fd == SOCKET_ERROR {
		return 0, ws2_32.WSAGetLastError()
	}

	return fd, 0
}

send(fd uint, buffer string) -> (count uint, err error) {
	n := ws2_32.send(fd, buffer.ptr, buffer.len as int, 0)

	if n == SOCKET_ERROR {
		return 0, ws2_32.WSAGetLastError()
	}

	return n as uint, 0
}

extern {
	ws2_32 {
		accept(fd uint, address *any|nil, length int) -> uint
		bind(fd uint, address *any, length int) -> int32
		closesocket(fd uint) -> int32
		connect(fd uint, address *any, length int) -> int32
		listen(fd uint, backlog int) -> int32
		recv(fd uint, address *byte|nil, length int, flags int) -> int32
		socket(family int, type int, protocol int) -> uint
		send(fd uint, address *byte, length int, flags int) -> int32
		shutdown(fd uint, how int) -> int32
		WSAGetLastError() -> int32
	}
}