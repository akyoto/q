package macho_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/config"
	"git.urbach.dev/cli/q/src/exe"
	"git.urbach.dev/cli/q/src/macho"
)

func TestWrite(t *testing.T) {
	macho.Write(&exe.Discard{}, &config.Build{Arch: config.ARM}, nil, nil)
	macho.Write(&exe.Discard{}, &config.Build{Arch: config.X86}, nil, nil)
	macho.Write(&exe.Discard{}, &config.Build{Arch: config.UnknownArch}, nil, nil)
}