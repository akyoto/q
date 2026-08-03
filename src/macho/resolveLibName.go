package macho

// resolveLibName returns the real libc library name and also appends the '.dylib' suffix as well as a 0-byte.
func resolveLibName(libName string) string {
	if libName == "System" {
		libName = "System.B"
	}

	return "lib" + libName + ".dylib\000"
}