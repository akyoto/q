init() {
	kernel32.SetConsoleCP(utf8)
	kernel32.SetConsoleOutputCP(utf8)
	stdin = kernel32.GetStdHandle(-10)
	stdout = kernel32.GetStdHandle(-11)
}

global {
	stdin int64
	stdout int64
}

const {
	utf8 = 65001
}

extern {
	kernel32 {
		GetStdHandle(device int64) -> (handle int64)
		SetConsoleCP(cp uint)
		SetConsoleOutputCP(cp uint)
	}
}