package codegen

import (
	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/config"
	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/cli/q/src/token"
)

// conditionalSet sets the target register to 0 or 1 depending on the condition.
func (f *Function) conditionalSet(register cpu.Register, op token.Kind, unsigned bool) {
	if f.build.Arch == config.X86 {
		f.Assembler.Append(&asm.MoveNumber{
			Destination: register,
			Number:      0,
		})
	}

	f.Assembler.Append(&asm.ConditionalSet{
		Destination: register,
		Condition:   tokenToCondition(op, unsigned),
	})
}