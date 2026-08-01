package macho

import (
	"git.urbach.dev/cli/q/src/dll"
	"git.urbach.dev/cli/q/src/exe"
)

const (
	BaseAddress  = 0x100000000
	NumSegments  = 4
	HashPageSize = 4096
	numCommands  = 10
	sizeCommands = NumSegments*Segment64Size +
		Section64Size +
		UuidSize +
		MainSize +
		BuildVersionSize +
		DylinkerCommandSize + len(LinkerString) +
		LinkEditDataCommandSize*2
)

// HeaderEnd returns the end of the headers.
func HeaderEnd(libs dll.List) int {
	return HeaderSize + SizeCommands(libs)
}

// NumCommands returns the number of load commands.
func NumCommands(libs dll.List) int {
	return numCommands + libs.Count()
}

// SizeCommands returns the size of load commands.
func SizeCommands(libs dll.List) int {
	size := sizeCommands

	for lib := range libs.All() {
		size += DylibCommandSize
		size += exe.Align(len(lib.Name)+1, 8)
	}

	return size
}