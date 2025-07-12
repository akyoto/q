package pe_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/config"
	"git.urbach.dev/cli/q/src/dll"
	"git.urbach.dev/cli/q/src/exe"
	"git.urbach.dev/cli/q/src/pe"
)

func TestWrite(t *testing.T) {
	libs := dll.List{}
	libs.Append("kernel32", "ExitProcess")
	libs.Append("user32", "MessageBox")

	pe.Write(&exe.Discard{}, &config.Build{Arch: config.ARM}, nil, nil, libs)
	pe.Write(&exe.Discard{}, &config.Build{Arch: config.X86}, nil, nil, libs)
	pe.Write(&exe.Discard{}, &config.Build{Arch: config.UnknownArch}, nil, nil, libs)
}