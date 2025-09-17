import io
import mem

args() -> !string {
	argcp := new(int32)
	wargv := shell32.CommandLineToArgvW(kernel32.GetCommandLineW(), argcp)
	argc := [argcp] as uint
	delete(argcp)
	count := 0

	loop i := 0..argc {
		count += kernel32.WideCharToMultiByte(io.utf8, 0, wargv[i], -1, 0, 0, 0, 0)
	}

	argv := mem.alloc(count)
	pos := 0

	loop i := 0..argc {
		pos += kernel32.WideCharToMultiByte(io.utf8, 0, wargv[i], -1, argv.ptr + pos, count - pos, 0, 0)
	}

	return argv
}

extern {
	kernel32 {
		GetCommandLineW() -> *uint16
		WideCharToMultiByte(codePage uint32, flags uint32, wideCharStr *uint16, wideChar int32, multiByteStr *byte|nil, multiByte int32, defaultChar *byte|nil, usedDefaultChar *int32|nil) -> (written int32)
	}

	shell32 {
		CommandLineToArgvW(cmdLine *uint16, numArgs *int32) -> **uint16
	}
}