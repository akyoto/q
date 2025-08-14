package codegen

import (
	"fmt"
	"slices"

	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/cli/q/src/types"
)

// exec executes a step which appends it to the assembler's instruction list.
func (f *Function) exec(step *Step) {
	if f.isLastStepInBlock(step) {
		for _, live := range step.Live {
			if live.Phi != nil && live.Register != live.Phi.Register {
				f.Assembler.Append(&asm.Move{
					Destination: live.Phi.Register,
					Source:      live.Register,
				})
			}
		}
	}

	switch instr := step.Value.(type) {
	case *ssa.Assert:
		f.jumpIfFalse(instr.Condition.(*ssa.BinaryOp).Op, "run.crash")

	case *ssa.BinaryOp:
		left := f.ValueToStep[instr.Left]
		right := f.ValueToStep[instr.Right]

		if instr.Op.IsComparison() {
			f.Assembler.Append(&asm.Compare{Destination: left.Register, Source: right.Register})
			return
		}

		switch instr.Op {
		case token.Add:
			f.Assembler.Append(&asm.Add{
				Destination: step.Register,
				Source:      left.Register,
				Operand:     right.Register,
			})

		case token.Div:
			f.Assembler.Append(&asm.Divide{
				Destination: step.Register,
				Source:      left.Register,
				Operand:     right.Register,
			})

		case token.Mul:
			f.Assembler.Append(&asm.Multiply{
				Destination: step.Register,
				Source:      left.Register,
				Operand:     right.Register,
			})

		case token.Sub:
			f.Assembler.Append(&asm.Subtract{
				Destination: step.Register,
				Source:      left.Register,
				Operand:     right.Register,
			})

		case token.Mod:
			f.Assembler.Append(&asm.Modulo{
				Destination: step.Register,
				Source:      left.Register,
				Operand:     right.Register,
			})

		case token.And:
			f.Assembler.Append(&asm.And{
				Destination: step.Register,
				Source:      left.Register,
				Operand:     right.Register,
			})

		case token.Or:
			f.Assembler.Append(&asm.Or{
				Destination: step.Register,
				Source:      left.Register,
				Operand:     right.Register,
			})

		case token.Xor:
			f.Assembler.Append(&asm.Xor{
				Destination: step.Register,
				Source:      left.Register,
				Operand:     right.Register,
			})

		case token.Shl:
			f.Assembler.Append(&asm.ShiftLeft{
				Destination: step.Register,
				Source:      left.Register,
				Operand:     right.Register,
			})

		case token.Shr:
			f.Assembler.Append(&asm.ShiftRightSigned{
				Destination: step.Register,
				Source:      left.Register,
				Operand:     right.Register,
			})

		default:
			panic(fmt.Sprintf("not implemented: %d", instr.Op))
		}

	case *ssa.Bool:
		if step.Register == -1 {
			return
		}

		number := 0

		if instr.Bool {
			number = 1
		}

		f.Assembler.Append(&asm.MoveNumber{
			Destination: step.Register,
			Number:      number,
		})

	case *ssa.Branch:
		var op token.Kind
		binaryOp, isBinaryOp := instr.Condition.(*ssa.BinaryOp)

		if isBinaryOp && binaryOp.Op.IsComparison() {
			op = binaryOp.Op
		} else {
			op = token.NotEqual

			f.Assembler.Append(&asm.CompareNumber{
				Destination: f.ValueToStep[instr.Condition].Register,
				Number:      0,
			})
		}

		following := f.Steps[step.Index+1].Value.(*Label)

		switch following.Name {
		case instr.Then.Label:
			f.jumpIfFalse(op, instr.Else.Label)
		case instr.Else.Label:
			f.jumpIfTrue(op, instr.Then.Label)
		default:
			panic("branch instruction must be followed by the 'then' or 'else' block")
		}

	case *ssa.Bytes:
		f.Count.Data++
		label := f.CreateLabel("data", f.Count.Data)
		f.Assembler.SetData(label, instr.Bytes)

		f.Assembler.Append(&asm.MoveLabel{
			Destination: step.Register,
			Label:       label,
		})

	case *ssa.Call:
		args := instr.Arguments

		for i, arg := range args {
			if f.ValueToStep[arg].Register == f.CPU.Call.In[i] {
				continue
			}

			f.Assembler.Append(&asm.Move{
				Destination: f.CPU.Call.In[i],
				Source:      f.ValueToStep[arg].Register,
			})
		}

		f.Assembler.Append(&asm.Call{Label: instr.Func.String()})

		if step.Register == -1 || step.Register == f.CPU.Call.Out[0] {
			return
		}

		f.Assembler.Append(&asm.Move{
			Destination: step.Register,
			Source:      f.CPU.Call.Out[0],
		})

	case *ssa.CallExtern:
		args := instr.Arguments
		var pushed []cpu.Register

		for i, arg := range slices.Backward(args) {
			if i >= len(f.CPU.ExternCall.In) {
				pushed = append(pushed, f.ValueToStep[arg].Register)
				continue
			}

			if f.ValueToStep[arg].Register == f.CPU.ExternCall.In[i] {
				continue
			}

			f.Assembler.Append(&asm.Move{
				Destination: f.CPU.ExternCall.In[i],
				Source:      f.ValueToStep[arg].Register,
			})
		}

		if len(pushed) > 0 {
			f.Assembler.Append(&asm.Push{Registers: pushed})
		}

		f.Assembler.Append(&asm.CallExtern{Library: instr.Func.Package, Function: instr.Func.Name})

		if len(pushed) > 0 {
			f.Assembler.Append(&asm.Pop{Registers: pushed})
		}

		if step.Register == -1 || step.Register == f.CPU.ExternCall.Out[0] {
			return
		}

		f.Assembler.Append(&asm.Move{
			Destination: step.Register,
			Source:      f.CPU.ExternCall.Out[0],
		})

	case *ssa.Int:
		if step.Register == -1 {
			return
		}

		f.Assembler.Append(&asm.MoveNumber{
			Destination: step.Register,
			Number:      instr.Int,
		})

	case *Label:
		f.Assembler.Append(&asm.Label{
			Name: instr.Name,
		})

	case *ssa.Load:
		if step.Register == -1 {
			return
		}

		address := f.ValueToStep[instr.Address]
		index := f.ValueToStep[instr.Index]
		elementSize := step.Value.Type().Size()

		f.Assembler.Append(&asm.Load{
			Base:        address.Register,
			Index:       index.Register,
			Destination: step.Register,
			Length:      byte(elementSize),
		})

	case *ssa.Parameter:
		source := f.CPU.Call.In[instr.Index]

		if step.Register == -1 || step.Register == source {
			return
		}

		f.Assembler.Append(&asm.Move{
			Destination: step.Register,
			Source:      source,
		})

	case *ssa.Jump:
		f.Assembler.Append(&asm.Jump{Label: instr.To.Label})

	case *ssa.Phi:
		// Phi does not generate any machine instructions.

	case *ssa.Return:
		defer f.leave()

		if len(instr.Arguments) == 0 {
			return
		}

		retVal := f.ValueToStep[instr.Arguments[0]]

		if retVal.Register == -1 || retVal.Register == f.CPU.Call.Out[0] {
			return
		}

		f.Assembler.Append(&asm.Move{
			Destination: f.CPU.Call.Out[0],
			Source:      retVal.Register,
		})

	case *ssa.Store:
		address := f.ValueToStep[instr.Address]
		index := f.ValueToStep[instr.Index]
		source := f.ValueToStep[instr.Value]
		pointer := address.Value.Type().(*types.Pointer)
		elementSize := pointer.To.Size()

		f.Assembler.Append(&asm.Store{
			Base:   address.Register,
			Index:  index.Register,
			Source: source.Register,
			Length: byte(elementSize),
		})

	case *ssa.Syscall:
		for i, arg := range instr.Arguments {
			if f.ValueToStep[arg].Register != f.CPU.Syscall.In[i] {
				f.Assembler.Append(&asm.Move{
					Destination: f.CPU.Syscall.In[i],
					Source:      f.ValueToStep[arg].Register,
				})
			}
		}

		f.Assembler.Append(&asm.Syscall{})

		if step.Register == -1 || step.Register == f.CPU.Syscall.Out[0] {
			return
		}

		f.Assembler.Append(&asm.Move{
			Destination: step.Register,
			Source:      f.CPU.Syscall.Out[0],
		})

	case *ssa.UnaryOp:
		left := f.ValueToStep[instr.Operand]

		switch instr.Op {
		case token.Negate:
			f.Assembler.Append(&asm.Negate{
				Destination: step.Register,
				Source:      left.Register,
			})
		}

	default:
		panic("not implemented: " + instr.String())
	}
}