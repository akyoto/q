import mem
import strings

run(path string) -> error {
	si := new(StartupInfo)
	si.size = 104
	pi := new(ProcessInformation)
	cpath := strings.c(path)
	success := kernel32.CreateProcessA(0, cpath.ptr, 0, 0, false, 0, 0, 0, si, pi)
	mem.free(cpath)

	if success == false {
		return -1
	}

	kernel32.WaitForSingleObject(pi.process, 0xFFFFFFFF)
	kernel32.CloseHandle(pi.process)
	kernel32.CloseHandle(pi.thread)
	return 0
}

extern {
	kernel32 {
		CloseHandle(handle *any) -> bool
		CreateProcessA(applicationName *byte, commandLine *byte, processAttributes *any, threadAttributes *any, inheritHandles bool, creationFlags uint32, environment *uint16, currentDirectory *uint16, startupInfo *any, processInformation *any) -> (success bool)
		WaitForSingleObject(handle *any, milliseconds uint32) -> uint32
	}
}