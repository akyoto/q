init() {
	wsa := new(WsaData)
	ws2_32.WSAStartup(0x0202, wsa)
}

exit() {
	ws2_32.WSACleanup()
}

extern {
	ws2_32 {
		WSAStartup(version uint16, address *WsaData)
		WSACleanup()
	}
}

WsaData {
	version uint16
	highVersion uint16
	maxSockets uint16
	maxUdpDg uint16
	vendorInfo *byte

	// TODO: 257 bytes (description)
	_ uint64
	_ uint64
	_ uint64
	_ uint64
	_ uint64
	_ uint64
	_ uint64
	_ uint64

	_ uint64
	_ uint64
	_ uint64
	_ uint64
	_ uint64
	_ uint64
	_ uint64
	_ uint64

	_ uint64
	_ uint64
	_ uint64
	_ uint64
	_ uint64
	_ uint64
	_ uint64
	_ uint64

	_ uint64
	_ uint64
	_ uint64
	_ uint64
	_ uint64
	_ uint64
	_ uint64
	_ uint64

	_ uint8

	// TODO: 129 bytes (systemStatus)
	_ uint64
	_ uint64
	_ uint64
	_ uint64
	_ uint64
	_ uint64
	_ uint64
	_ uint64

	_ uint64
	_ uint64
	_ uint64
	_ uint64
	_ uint64
	_ uint64
	_ uint64
	_ uint64

	_ uint8

	// TODO: 6 bytes (padding)
	_ uint32
	_ uint16
}