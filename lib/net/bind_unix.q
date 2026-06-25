import bits

bind(fd uint, port uint16) -> error {
	addr := new(AddressIPv6) {
		family: AF_INET6,
		port: bits.reverseBytes(port),
	}

	return syscall(_bind, fd, addr, 28)
}