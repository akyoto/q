exit(code int) {
	kernel32.ExitProcess(code)
}

extern {
	kernel32 {
		ExitProcess(code uint)
	}
}