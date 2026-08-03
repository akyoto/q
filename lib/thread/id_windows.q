id() -> uint {
	return kernel32.GetCurrentThreadId() as uint
}

extern {
	kernel32 {
		GetCurrentThreadId() -> int
	}
}