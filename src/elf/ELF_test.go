package elf_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/config"
	"git.urbach.dev/cli/q/src/elf"
	"git.urbach.dev/cli/q/src/exe"
)

func TestWrite(t *testing.T) {
	elf.Write(&exe.Discard{}, &config.Build{Arch: config.ARM}, nil, nil)
	elf.Write(&exe.Discard{}, &config.Build{Arch: config.X86}, nil, nil)
	elf.Write(&exe.Discard{}, &config.Build{Arch: config.UnknownArch}, nil, nil)
}