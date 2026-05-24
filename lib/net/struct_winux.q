AddressIPv4 {
	family uint16
	port uint16
	addr uint64
	_ uint64
}

AddressIPv6 {
	family uint16
	port uint16
	flowinfo uint32
	_ uint32
	_ uint32
	_ uint32
	_ uint32
	scope uint32
	_ uint32
}