read(buffer string) -> (count uint, err error) {
	return readFrom(stdin, buffer)
}

readFrom(fd uint, buffer string) -> (count uint, err error) {
	ptr := new(uint32)
	success := kernel32.ReadFile(fd, buffer.ptr, buffer.len as uint32, ptr, 0)

	if !success {
		return 0, -1
	}

	count := [ptr]
	return count as uint, 0
}