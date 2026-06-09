id() -> int {
	return kernel32.GetCurrentThreadId()
}

extern {
	kernel32 {
		GetCurrentThreadId() -> int
	}
}