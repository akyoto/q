import run

alloc(length uint) -> (buffer !string) {
	if length == 0 {
		return string{ptr: 0 as *uint8, len: 0}
	}

	x := kernel32.VirtualAlloc(0, length, commit|reserve, readwrite)

	if x == 0 {
		run.crash()
	}

	return string{ptr: x, len: length}
}

free(buffer !string) {
	if buffer.ptr == 0 {
		return
	}

	kernel32.VirtualFree(buffer.ptr, buffer.len, decommit)
}

extern {
	kernel32 {
		VirtualAlloc(address int, size uint, flags uint32, protection uint32) -> (address *byte)
		VirtualFree(address *any, size uint, type uint32) -> (success bool)
	}
}