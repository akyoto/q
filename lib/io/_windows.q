extern {
	kernel32 {
		ReadFile(fd uint, buffer *byte, length uint32, count *uint32, overlapped *any|nil) -> (success bool)
		WriteFile(fd uint, buffer *byte, length uint32, count *uint32, overlapped *any|nil) -> (success bool)
	}
}