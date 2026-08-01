package macho

// Dylib is a load command followed by the library name.
type Dylib struct {
	DylibCommand
	Name []byte
}