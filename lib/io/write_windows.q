write(buffer string) -> (written int) {
	ptr := new(uint32)
	kernel32.WriteFile(stdout, buffer.ptr, buffer.len as uint32, ptr, 0)
	count := [ptr]
	delete(ptr)
	return count as int
}

writeTo(fd int, buffer string) -> (written int, err error) {
	ptr := new(uint32)
	success := kernel32.WriteFile(fd, buffer.ptr, buffer.len as uint32, ptr, 0)

	if !success {
		return 0, -1
	}

	count := [ptr]
	delete(ptr)
	return count as int, 0
}

extern {
	kernel32 {
		WriteFile(fd int64, buffer *byte, length uint32, written *uint32, overlapped *any|nil) -> (success bool)
	}
}