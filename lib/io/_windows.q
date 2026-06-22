extern {
	kernel32 {
		ReadFile(fd int64, buffer *byte, length uint32, count *uint32, overlapped *any|nil) -> (success bool)
		WriteFile(fd int64, buffer *byte, length uint32, count *uint32, overlapped *any|nil) -> (success bool)
	}
}