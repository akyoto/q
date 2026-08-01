package macho

// resolveLibName returns the real libc library name and also appends the '.dylib' suffix as well as a 0-byte.
func resolveLibName(libName string) string {
	if libName == "libc" {
		libName = "libSystem.B"
	}

	return libName + ".dylib\000"
}