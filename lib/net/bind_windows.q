import bits

bind(fd uint, port uint16) -> error {
	addr := new(AddressIPv6) {
		family: AF_INET6,
		port: bits.reverseBytes(port),
	}

	err := ws2_32.bind(fd, addr, 28)

	if err == SOCKET_ERROR {
		return ws2_32.WSAGetLastError()
	}

	return 0
}