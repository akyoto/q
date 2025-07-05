write(buffer string) -> (written int) {
	stdout := kernel32.GetStdHandle(-11)
	kernel32.WriteFile(stdout, buffer.ptr, buffer.len, 0, 0)
	return buffer.len
}

extern {
	kernel32 {
		GetStdHandle(device int64) -> (handle int64)
		WriteFile(fd int64, buffer *byte, length uint32, written *uint32, overlapped *any) -> (success bool)
	}
}