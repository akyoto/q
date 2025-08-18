import os

init() {
	utf8 := 65001
	kernel32.SetConsoleCP(utf8)
	kernel32.SetConsoleOutputCP(utf8)
	wsa := new(WsaData)
	ws2_32.WSAStartup(0x0202, wsa)
	main.main()
	ws2_32.WSACleanup()
	os.exit(0)
}

crash() {
	os.exit(1)
}

WsaData {
	wVersion uint16
	wHighVersion uint16
	iMaxSockets uint16
	iMaxUdpDg uint16
	lpVendorInfo *byte
}

extern {
	kernel32 {
		SetConsoleCP(cp uint)
		SetConsoleOutputCP(cp uint)
	}

	ws2_32 {
		WSAStartup(version uint16, address *any)
		WSACleanup()
	}
}