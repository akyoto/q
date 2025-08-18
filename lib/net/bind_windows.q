sockaddr_in {
	sin_family int16
	sin_port uint16
	sin_addr int64
	sin_zero int64
}

bind(socket int, port uint16) -> (error int) {
	addr := new(sockaddr_in)
	addr.sin_family = 2
	addr.sin_port = htons(port)
	return ws2_32.bind(socket, addr, 16)
}