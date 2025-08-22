listenTCP(address int64, port uint16) -> int {
	s := socket(v4, tcp, 0)

	if s < 0 {
		return s
	}

	if bind(s, address, port) != 0 {
		return -1
	}

	if listen(s, backlog) != 0 {
		return -1
	}

	return s
}