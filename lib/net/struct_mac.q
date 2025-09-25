AddressIPv4 {
	len uint8
	family uint8
	port uint16
	addr uint64
	_ uint64
}

AddressIPv6 {
	len uint8
	family uint8
	port uint16
	flowinfo uint32
	_ uint32
	_ uint32
	_ uint32
	_ uint32
	scope uint32
	_ uint32
}