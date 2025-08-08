package macho

// Version creates the version for the build version load command.
func Version(major uint32, minor uint32, patch uint32) uint32 {
	return major<<16 | minor<<8 | patch
}