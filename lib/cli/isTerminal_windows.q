isTerminal(fd uint) -> bool {
	mode := new(uint32)
	return kernel32.GetConsoleMode(fd, mode)
}

extern {
	kernel32 {
		GetConsoleMode(fd uint, mode *uint32) -> (success bool)
	}
}