import run

rawAlloc(length uint) -> *uint8 {
	x := kernel32.VirtualAlloc(0, length, commit|reserve, readwrite)

	if x == 0 {
		run.crash()
	}

	return x
}

rawFree(ptr *any, len uint) {
	kernel32.VirtualFree(ptr, len, decommit)
}

extern {
	kernel32 {
		VirtualAlloc(address int, size uint, flags uint32, protection uint32) -> (address *byte)
		VirtualFree(address *any, size uint, type uint32) -> (success bool)
	}
}