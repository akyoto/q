import io
import mem

args() -> ![]string {
	argcp := new(int32)
	wargv := shell32.CommandLineToArgvW(kernel32.GetCommandLineW(), argcp)
	argc := [argcp]-1 as uint
	delete(argcp)
	args := new(string, argc)

	loop i := 0..args.len {
		length := kernel32.WideCharToMultiByte(io.utf8, 0, wargv[i+1], -1, 0, 0, 0, 0) - 1 as uint
		// TODO: Memory leak that needs to be removed once the type system has been fully implemented
		args[i] = string{ptr: mem.rawAlloc(length), len: length}
		kernel32.WideCharToMultiByte(io.utf8, 0, wargv[i+1], -1, args[i].ptr, length as int32, 0, 0)
	}

	kernel32.LocalFree(wargv)
	return args
}

extern {
	kernel32 {
		GetCommandLineW() -> *uint16
		LocalFree(mem *any) -> *any|nil
		WideCharToMultiByte(codePage uint32, flags uint32, wideCharStr *uint16, wideChar int32, multiByteStr *byte|nil, multiByte int32, defaultChar *byte|nil, usedDefaultChar *int32|nil) -> (written int32)
	}

	shell32 {
		CommandLineToArgvW(cmdLine *uint16, numArgs *int32) -> **uint16
	}
}