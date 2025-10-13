read(buffer string) -> (read int) {
	ptr := new(uint32)
	kernel32.ReadFile(stdin, buffer.ptr, buffer.len as uint32, ptr, 0)
	count := [ptr]
	delete(ptr)
	return count as int
}

readFrom(fd int, buffer string) -> (read int, err error) {
	ptr := new(uint32)
	success := kernel32.ReadFile(fd, buffer.ptr, buffer.len as uint32, ptr, 0)

	if !success {
		return 0, -1
	}

	count := [ptr]
	delete(ptr)
	return count as int, 0
}

extern {
	kernel32 {
		ReadFile(fd int64, buffer *byte, length uint32, read *uint32, overlapped *any|nil) -> (success bool)
	}
}