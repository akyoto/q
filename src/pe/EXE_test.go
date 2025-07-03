package pe_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/build"
	"git.urbach.dev/cli/q/src/exe"
	"git.urbach.dev/cli/q/src/pe"
)

func TestWrite(t *testing.T) {
	pe.Write(&exe.Discard{}, &build.Build{Arch: build.ARM}, nil, nil, nil)
	pe.Write(&exe.Discard{}, &build.Build{Arch: build.X86}, nil, nil, nil)
	pe.Write(&exe.Discard{}, &build.Build{Arch: build.UnknownArch}, nil, nil, nil)
}