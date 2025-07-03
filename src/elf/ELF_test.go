package elf_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/build"
	"git.urbach.dev/cli/q/src/elf"
	"git.urbach.dev/cli/q/src/exe"
)

func TestWrite(t *testing.T) {
	elf.Write(&exe.Discard{}, &build.Build{Arch: build.ARM}, nil, nil)
	elf.Write(&exe.Discard{}, &build.Build{Arch: build.X86}, nil, nil)
	elf.Write(&exe.Discard{}, &build.Build{Arch: build.UnknownArch}, nil, nil)
}