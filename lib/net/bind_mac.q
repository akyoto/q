sockaddr_in_bsd {
	sin_len int8
	sin_family int8
	sin_port uint16
	sin_addr int64
	sin_zero int64
}

bind(socket int, address int64, port uint16) -> (error int) {
	addr := new(sockaddr_in_bsd)
	addr.sin_family = 2
	addr.sin_port = htons(port)
	addr.sin_addr = address
	return syscall(_bind, socket, addr, 16)
}