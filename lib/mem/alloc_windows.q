import run

alloc(length uint) -> (buffer !string) {
	x := kernel32.VirtualAlloc(0, length, commit|reserve, readwrite)

	if x == 0 {
		run.crash()
	}

	return string{ptr: x, len: length}
}

free(buffer !string) {
	kernel32.VirtualFree(buffer.ptr, buffer.len, decommit)
}

extern {
	kernel32 {
		VirtualAlloc(address int, size uint, flags uint32, protection uint32) -> (address *byte)
		VirtualFree(address *any, size uint, type uint32) -> (success bool)
	}
}