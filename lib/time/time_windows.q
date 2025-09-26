now() -> int {
	systemTime := new(int64)
	kernel32.GetSystemTimePreciseAsFileTime(systemTime)
	n := [systemTime]
	delete(systemTime)
	return (n - 116444736000000000) * 100
}

sleep(nanoseconds int) {
	milliseconds := nanoseconds / millisecond
	kernel32.Sleep(milliseconds as uint32)
}

extern {
	kernel32 {
		Sleep(milliseconds uint32)
		GetSystemTimePreciseAsFileTime(systemTime *int64)
	}
}