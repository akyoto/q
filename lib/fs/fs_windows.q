openRead(path *byte) -> (!uint, error) {
	fd := kernel32.CreateFileA(path, GENERIC_READ, FILE_SHARE_READ, 0, OPEN_EXISTING, FILE_ATTRIBUTE_NORMAL, 0)

	if fd == INVALID_HANDLE_VALUE {
		return 0, kernel32.GetLastError()
	}

	return fd, 0
}

openWrite(path *byte) -> (!uint, error) {
	fd := kernel32.CreateFileA(path, GENERIC_WRITE, FILE_SHARE_WRITE, 0, CREATE_ALWAYS, FILE_ATTRIBUTE_NORMAL, 0)

	if fd == INVALID_HANDLE_VALUE {
		return 0, kernel32.GetLastError()
	}

	return fd, 0
}

size(fd uint) -> (uint, error) {
	ptr := new(int64)
	success := kernel32.GetFileSizeEx(fd, ptr)

	if success {
		size := [ptr] as uint
		return size, 0
	}

	return 0, -1
}

close(fd !uint) -> error {
	success := kernel32.CloseHandle(fd)

	if success {
		return 0
	}

	return -1
}

memfd_create(_path *byte, _flags uint) -> (!uint, error) {
	return 0, -1
}

ftruncate(_fd uint, _length uint) -> error {
	return -1
}

extern {
	kernel32 {
		CloseHandle(fd uint) -> (success bool)
		CreateFileA(path *byte, desiredAccess uint32, shareMode uint32, securityAttributes *any|nil, creationDisposition uint32, flagsAndAttributes uint32, templateFile int) -> (fd uint)
		GetFileSizeEx(fd uint, fileSize *int64) -> (success bool)
		GetLastError() -> uint32
	}
}