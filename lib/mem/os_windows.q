import run

osAlloc(length uint) -> *uint8 {
	if length == 0 {
		return 0 as *uint8
	}

	x := kernel32.VirtualAlloc(0, length, commit|reserve, readwrite)

	if x == 0 {
		run.crash()
	}

	return x
}

osFree(ptr *any, len uint) {
	if ptr == 0 {
		return
	}

	kernel32.VirtualFree(ptr, len, decommit)
}

extern {
	kernel32 {
		VirtualAlloc(address int, size uint, flags uint32, protection uint32) -> (address *byte)
		VirtualFree(address *any, size uint, type uint32) -> (success bool)
	}
}