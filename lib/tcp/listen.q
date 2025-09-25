import net

listen(port uint16) -> (int, error) {
	s, err := net.socket(net.v6, tcp, 0)

	if err != 0 {
		return 0, err
	}

	err := net.bind(s, port)

	if err != 0 {
		return 0, err
	}

	err := net.listen(s, backlog)

	if err != 0 {
		return 0, err
	}

	return s, 0
}