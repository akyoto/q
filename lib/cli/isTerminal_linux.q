isTerminal(fd uint) -> bool {
	settings := new(TerminalIOSettings)
	return syscall(_ioctl, fd, TCGETS2, settings) == 0
}