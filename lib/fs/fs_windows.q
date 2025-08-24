open(path *byte, _flags int, _mode int) -> int {
	return kernel32.CreateFileA(path, GENERIC_READ, FILE_SHARE_READ, 0, OPEN_EXISTING, FILE_ATTRIBUTE_NORMAL, 0)
}

size(fd int) -> int {
	fileSize := new(int64)
	kernel32.GetFileSizeEx(fd, fileSize)
	return [fileSize]
}

close(fd int) -> bool {
	return kernel32.CloseHandle(fd)
}

extern {
	kernel32 {
		CloseHandle(handle int) -> (success bool)
		CreateFileA(path *byte, desiredAccess uint32, shareMode uint32, securityAttributes *any, creationDisposition uint32, flagsAndAttributes uint32, templateFile int) -> (fd int)
		GetFileSizeEx(fd int, fileSize *int64) -> (success bool)
	}
}