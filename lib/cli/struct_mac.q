TerminalIOSettings {
	c_iflag uint64
	c_oflag uint64
	c_cflag uint64
	c_lflag uint64

	// TODO: 20 bytes (c_cc)
	_ uint64
	_ uint64
	_ uint16
	_ uint16

	// Padding
	_ uint32

	c_ispeed uint64
	c_ospeed uint64
}