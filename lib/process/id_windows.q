id() -> int {
	return kernel32.GetCurrentProcessId()
}

extern {
	kernel32 {
		GetCurrentProcessId() -> int
	}
}