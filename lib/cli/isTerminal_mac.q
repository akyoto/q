isTerminal(fd uint) -> bool {
	settings := new(TerminalIOSettings)
	return syscall(_ioctl, fd, TIOCGETA, settings) == 0
}