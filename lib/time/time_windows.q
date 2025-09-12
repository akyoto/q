sleep(nanoseconds int) {
	milliseconds := nanoseconds / 1000000
	kernel32.Sleep(milliseconds as uint32)
}

extern {
	kernel32 {
		Sleep(milliseconds uint32)
	}
}