write(buffer string) -> (count uint, err error) {
	return writeTo(stdout, buffer)
}

writeTo(fd int, buffer string) -> (count uint, err error) {
	ptr := new(uint32)
	success := kernel32.WriteFile(fd, buffer.ptr, buffer.len as uint32, ptr, 0)

	if !success {
		return 0, -1
	}

	count := [ptr]
	return count as uint, 0
}