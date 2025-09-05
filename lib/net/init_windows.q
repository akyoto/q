init() {
	wsa := new(WsaData)
	ws2_32.WSAStartup(0x0202, wsa)
	delete(wsa)
}

exit() {
	ws2_32.WSACleanup()
}

extern {
	ws2_32 {
		WSAStartup(version uint16, address *any)
		WSACleanup()
	}
}

WsaData {
	version uint16
	highVersion uint16
	maxSockets uint16
	maxUdpDg uint16
	vendorInfo *byte
}