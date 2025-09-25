import math

bind(socket int, port uint16) -> error {
	addr := new(AddressIPv6)
	addr.family = v6
	addr.port = math.reverseBytes(port)
	err := ws2_32.bind(socket, addr, 28)
	delete(addr)
	return err
}