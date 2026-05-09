init() {
	envp = kernel32.GetEnvironmentStrings()
}

extern {
	kernel32 {
		GetEnvironmentStrings() -> *byte
	}
}