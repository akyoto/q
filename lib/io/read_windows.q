read(buffer string) -> (count uint, err error) {
	return readFrom(stdin, buffer)
}

readFrom(fd int, buffer string) -> (count uint, err error) {
	ptr := new(uint32)
	success := kernel32.ReadFile(fd, buffer.ptr, buffer.len as uint32, ptr, 0)

	if !success {
		delete(ptr)
		return 0, -1
	}

	count := [ptr]
	delete(ptr)
	return count as uint, 0
}

extern {
	kernel32 {
		ReadFile(fd int64, buffer *byte, length uint32, read *uint32, overlapped *any|nil) -> (success bool)
	}
}