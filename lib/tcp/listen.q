import net

listen(_address string, port uint16) -> int {
	s := net.socket(v4, tcp, 0)

	if s < 0 {
		return s
	}

	if net.bind(s, 0, port) != 0 {
		return -1
	}

	if net.listen(s, backlog) != 0 {
		return -1
	}

	return s
}