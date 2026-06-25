import net

listen(port uint16) -> (uint, error) {
	s, err := net.socket(net.AF_INET6, net.SOCK_STREAM, 0)

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