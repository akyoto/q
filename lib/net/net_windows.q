accept(fd int, address *any, length int) -> int {
	return ws2_32.accept(fd, address, length)
}

close(fd int) -> int {
	return ws2_32.shutdown(fd, 1)
	//return ws2_32.closesocket(fd)
}

listen(fd int, backlog int) -> int {
	return ws2_32.listen(fd, backlog)
}

recv(fd int, buffer string) -> int {
	return ws2_32.recv(fd, buffer.ptr, buffer.len, 0)
}

socket(family int, type int, protocol int) -> int {
	return ws2_32.socket(family, type, protocol)
}

send(fd int, buffer string) -> int {
	return ws2_32.send(fd, buffer.ptr, buffer.len, 0)
}

extern {
	ws2_32 {
		accept(fd int, address *any, length int) -> int
		bind(socket int, address *any, length int) -> (error int)
		closesocket(fd int) -> int
		listen(fd int, backlog int) -> int
		recv(fd int, address *byte, length int, flags int) -> int
		socket(family int, type int, protocol int) -> int
		send(fd int, address *byte, length int, flags int) -> int
		shutdown(fd int, how int) -> int
	}
}