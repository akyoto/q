package pe_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/config"
	"git.urbach.dev/cli/q/src/exe"
	"git.urbach.dev/cli/q/src/pe"
)

func TestWrite(t *testing.T) {
	pe.Write(&exe.Discard{}, &config.Build{Arch: config.ARM}, nil, nil, nil)
	pe.Write(&exe.Discard{}, &config.Build{Arch: config.X86}, nil, nil, nil)
	pe.Write(&exe.Discard{}, &config.Build{Arch: config.UnknownArch}, nil, nil, nil)
}