import math

sockaddr_in {
	sin_family int16
	sin_port uint16
	sin_addr int64
	sin_zero int64
}

bind(socket int, address int64, port uint16) -> error {
	addr := new(sockaddr_in)
	addr.sin_family = 2
	addr.sin_port = math.reverseBytes(port)
	addr.sin_addr = address
	err := syscall(_bind, socket, addr, 16)
	delete(addr)
	return err
}