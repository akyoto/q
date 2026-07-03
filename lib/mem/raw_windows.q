import run

rawAlloc(length uint) -> *uint8 {
	x := kernel32.VirtualAlloc(0, length, MEM_COMMIT|MEM_RESERVE, PAGE_READWRITE)

	if x == 0 {
		run.crash()
	}

	return x
}

rawFree(ptr *any, _len uint) {
	kernel32.VirtualFree(ptr, 0, MEM_RELEASE)
}

extern {
	kernel32 {
		VirtualAlloc(address int, size uint, flags uint32, protection uint32) -> (address *byte)
		VirtualFree(address *any, size uint, type uint32) -> (success bool)
	}
}