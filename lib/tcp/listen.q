import net

listen(_address string, port uint16) -> (int, error) {
	s, err := net.socket(v4, tcp, 0)

	if err != 0 {
		return 0, err
	}

	err := net.bind(s, 0, port)

	if err != 0 {
		return 0, err
	}

	err := net.listen(s, backlog)

	if err != 0 {
		return 0, err
	}

	return s, 0
}