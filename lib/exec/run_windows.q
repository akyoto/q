import mem
import strings

run(path string) -> error {
	startInfo := new(StartupInfo)
	startInfo.size = 104
	processInfo := new(ProcessInformation)
	cpath := strings.c(path)
	success := kernel32.CreateProcessA(0, cpath.ptr, 0, 0, false, 0, 0, 0, startInfo, processInfo)
	mem.free(cpath)

	if success == false {
		delete(processInfo)
		delete(startInfo)
		return -1
	}

	kernel32.WaitForSingleObject(processInfo.process, 0xFFFFFFFF)
	kernel32.CloseHandle(processInfo.process)
	kernel32.CloseHandle(processInfo.thread)
	delete(processInfo)
	delete(startInfo)
	return 0
}

extern {
	kernel32 {
		CloseHandle(handle *any) -> bool
		CreateProcessA(applicationName *byte, commandLine *byte, processAttributes *any, threadAttributes *any, inheritHandles bool, creationFlags uint32, environment *uint16, currentDirectory *uint16, startupInfo *any, processInformation *any) -> (success bool)
		WaitForSingleObject(handle *any, milliseconds uint32) -> uint32
	}
}