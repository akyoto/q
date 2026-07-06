TerminalIOSettings {
	c_iflag uint32
	c_oflag uint32
	c_cflag uint32
	c_lflag uint32
	c_line uint8

	// TODO: 19 bytes (c_cc)
	_ uint64
	_ uint64
	_ uint16
	_ uint8

	c_ispeed uint32
	c_ospeed uint32
}