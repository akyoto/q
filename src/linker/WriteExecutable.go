package linker

import (
	"encoding/binary"
	"os"

	"git.urbach.dev/cli/q/src/arm"
	"git.urbach.dev/cli/q/src/build"
	"git.urbach.dev/cli/q/src/core"
	"git.urbach.dev/cli/q/src/elf"
	"git.urbach.dev/cli/q/src/x86"
)

// WriteExecutable writes an executable file to disk.
func WriteExecutable(b *build.Build, result *core.Environment) error {
	executable := b.Executable()
	file, err := os.Create(executable)

	if err != nil {
		return err
	}

	code := []byte{}
	data := []byte{}

	switch b.Arch {
	case build.ARM:
		code = arm.MoveRegisterNumber(code, arm.X8, 93)
		code = arm.MoveRegisterNumber(code, arm.X0, 0)
		code = binary.LittleEndian.AppendUint32(code, arm.Syscall())

	case build.X86:
		code = x86.MoveRegisterNumber(code, x86.R0, 60)
		code = x86.MoveRegisterNumber(code, x86.R7, 0)
		code = x86.Syscall(code)
	}

	switch b.OS {
	case build.Linux:
		elf.Write(file, b, code, data)
	}

	err = file.Close()

	if err != nil {
		return err
	}

	return os.Chmod(executable, 0755)
}