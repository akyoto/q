import run

alloc(length int) -> (address *byte) {
	x := kernel32.VirtualAlloc(0, length, commit|reserve, readwrite)

	if x == 0 {
		run.crash()
	}

	return x
}

free(address *any, length int) {
	kernel32.VirtualFree(address, length, decommit)
}

extern {
	kernel32 {
		VirtualAlloc(address int, size uint, flags uint32, protection uint32) -> (address *byte)
		VirtualFree(address *any, size uint, type uint32) -> (success bool)
	}
}