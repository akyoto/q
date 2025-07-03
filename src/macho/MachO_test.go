package macho_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/build"
	"git.urbach.dev/cli/q/src/exe"
	"git.urbach.dev/cli/q/src/macho"
)

func TestWrite(t *testing.T) {
	macho.Write(&exe.Discard{}, &build.Build{Arch: build.ARM}, nil, nil)
	macho.Write(&exe.Discard{}, &build.Build{Arch: build.X86}, nil, nil)
	macho.Write(&exe.Discard{}, &build.Build{Arch: build.UnknownArch}, nil, nil)
}