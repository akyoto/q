import bits

bind(socket int, port uint16) -> error {
	addr := new(AddressIPv6) {
		family: AF_INET6,
		port: bits.reverseBytes(port),
	}

	err := syscall(_bind, socket, addr, 28)
	return err
}