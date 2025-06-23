package elf_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/build"
	"git.urbach.dev/cli/q/src/elf"
	"git.urbach.dev/cli/q/src/exe"
)

func TestWrite(t *testing.T) {
	elf.Write(&exe.Discard{}, &build.Build{Arch: build.ARM, FileAlign: 0x4000, MemoryAlign: 0x4000}, nil, nil)
	elf.Write(&exe.Discard{}, &build.Build{Arch: build.X86, FileAlign: 0x1000, MemoryAlign: 0x1000}, nil, nil)
	elf.Write(&exe.Discard{}, &build.Build{Arch: build.UnknownArch, FileAlign: 0x1000, MemoryAlign: 0x1000}, nil, nil)
}