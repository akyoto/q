TerminalIOSettings {
	c_iflag uint64
	c_oflag uint64
	c_cflag uint64
	c_lflag uint64
	c_cc [20]byte
	_ uint32
	c_ispeed uint64
	c_ospeed uint64
}