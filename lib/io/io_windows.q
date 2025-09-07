read(buffer string) -> (read int) {
	stdin := kernel32.GetStdHandle(-10)
	ptr := new(uint32)
	kernel32.ReadFile(stdin, buffer.ptr, buffer.len, ptr, 0)
	count := [ptr]
	delete(ptr)
	return count
}

readFrom(fd int, buffer string) -> (read int) {
	ptr := new(uint32)
	kernel32.ReadFile(fd, buffer.ptr, buffer.len, ptr, 0)
	count := [ptr]
	delete(ptr)
	return count
}

write(buffer string) -> (written int) {
	stdout := kernel32.GetStdHandle(-11)
	ptr := new(uint32)
	kernel32.WriteFile(stdout, buffer.ptr, buffer.len, ptr, 0)
	count := [ptr]
	delete(ptr)
	return count
}

writeTo(fd int, buffer string) -> (written int) {
	ptr := new(uint32)
	kernel32.WriteFile(fd, buffer.ptr, buffer.len, ptr, 0)
	count := [ptr]
	delete(ptr)
	return count
}

extern {
	kernel32 {
		GetStdHandle(device int64) -> (handle int64)
		ReadFile(fd int64, buffer *byte, length uint32, read *uint32, overlapped *any|nil) -> (success bool)
		WriteFile(fd int64, buffer *byte, length uint32, written *uint32, overlapped *any|nil) -> (success bool)
	}
}