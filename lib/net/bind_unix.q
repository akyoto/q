import bits

bind(socket int, port uint16) -> error {
	addr := new(AddressIPv6)
	addr.family = AF_INET6
	addr.port = bits.reverseBytes(port)
	err := syscall(_bind, socket, addr, 28)
	delete(addr)
	return err
}