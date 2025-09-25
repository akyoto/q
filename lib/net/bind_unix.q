import math

bind(socket int, address int64, port uint16) -> error {
	addr := new(AddressIPv4)
	addr.family = 2
	addr.port = math.reverseBytes(port)
	addr.addr = address
	err := syscall(_bind, socket, addr, 16)
	delete(addr)
	return err
}