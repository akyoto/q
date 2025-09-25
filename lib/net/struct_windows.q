AddressIPv4 {
	family int16
	port uint16
	addr int64
	zero int64
}

WsaData {
	version uint16
	highVersion uint16
	maxSockets uint16
	maxUdpDg uint16
	vendorInfo *byte
}