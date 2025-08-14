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

const {
	readwrite 0x0004
	commit 0x1000
	reserve 0x2000
	decommit 0x4000
}

extern {
	kernel32 {
		VirtualAlloc(address int, size uint, flags uint32, protection uint32) -> (address *byte)
		VirtualFree(address *any, size uint, type uint32) -> (success bool)
	}
}