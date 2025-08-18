read(buffer string) -> (read int) {
	stdin := kernel32.GetStdHandle(-10)
	count := new(uint32)
	kernel32.ReadFile(stdin, buffer.ptr, buffer.len, count, 0)
	return [count]
}

readFrom(fd int, buffer string) -> (read int) {
	count := new(uint32)
	kernel32.ReadFile(fd, buffer.ptr, buffer.len, count, 0)
	return [count]
}

write(buffer string) -> (written int) {
	stdout := kernel32.GetStdHandle(-11)
	count := new(uint32)
	kernel32.WriteFile(stdout, buffer.ptr, buffer.len, count, 0)
	return [count]
}

writeTo(fd int, buffer string) -> (written int) {
	count := new(uint32)
	kernel32.WriteFile(fd, buffer.ptr, buffer.len, count, 0)
	return [count]
}

extern {
	kernel32 {
		GetStdHandle(device int64) -> (handle int64)
		ReadFile(fd int64, buffer *byte, length uint32, read *uint32, overlapped *any) -> (success bool)
		WriteFile(fd int64, buffer *byte, length uint32, written *uint32, overlapped *any) -> (success bool)
	}
}