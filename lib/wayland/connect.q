import unix

connect() -> (!uint, error) {
	path, err := path()

	if err != 0 {
		return 0, err
	}

	socket, err := unix.connect(path)

	if err != 0 {
		return 0, err
	}

	return socket, 0
}