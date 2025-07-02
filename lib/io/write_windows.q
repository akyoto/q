write(_ string) -> (written int) {
	return 0
}

extern {
	kernel32 {
		GetStdHandle(handle int64) -> int64
		WriteConsoleA(fd int64, buffer *byte, length uint32, written *uint32) -> bool
	}
}