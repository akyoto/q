open(path *byte, _flags int, _mode int) -> (!int, error) {
	fd := kernel32.CreateFileA(path, GENERIC_READ, FILE_SHARE_READ, 0, OPEN_EXISTING, FILE_ATTRIBUTE_NORMAL, 0)

	if fd == -1 {
		return 0, fd
	}

	return fd, 0
}

size(fd int) -> (uint, error) {
	ptr := new(int64)
	success := kernel32.GetFileSizeEx(fd, ptr)

	if success {
		size := [ptr] as uint
		delete(ptr)
		return size, 0
	}

	delete(ptr)
	return 0, -1
}

close(fd !int) -> error {
	success := kernel32.CloseHandle(fd)

	if success {
		return 0
	}

	return -1
}

extern {
	kernel32 {
		CloseHandle(handle int) -> (success bool)
		CreateFileA(path *byte, desiredAccess uint32, shareMode uint32, securityAttributes *any|nil, creationDisposition uint32, flagsAndAttributes uint32, templateFile int) -> (fd int)
		GetFileSizeEx(fd int, fileSize *int64) -> (success bool)
	}
}