package macho

const BuildVersionSize = 24

// BuildVersion is the structure of the LC_BUILD_VERSION load command.
type BuildVersion struct {
	LoadCommand
	Length   uint32
	Platform Platform
	MinOS    uint32
	Sdk      uint32
	NumTools uint32
}