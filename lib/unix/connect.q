import net

connect(path string) -> (!int, error) {
	if path.len > maxPathLength {
		return 0, -1
	}

	s, err := net.socket(net.AF_UNIX, net.SOCK_STREAM, 0)

	if err != 0 {
		return 0, err
	}

	addr := address(path)
	err := net.connect(s, addr.ptr, addr.len)
	delete(addr)
	return s, err
}