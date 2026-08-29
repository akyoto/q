package verbose

import (
	"fmt"
	"reflect"

	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/core"
	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/go/color/ansi"
)

// ASM shows the assembly code.
func ASM(env *core.Environment, pattern string) {
	for f := range env.LiveFunctions() {
		if filter(f.FullName, pattern) {
			continue
		}

		printAssembly(f)
	}
}

// printAssembly shows the assembly instructions.
func printAssembly(f *core.Function) {
	var (
		mnemonic = ansi.Green
		function = ansi.Yellow
		imm      = ansi.Cyan
		label    = ansi.Reset
		other    = ansi.Reset
		register = ansi.Reset
	)

	rrr := func(name string, d cpu.Register, s cpu.Register, o cpu.Register) {
		mnemonic.Print("  " + name)
		register.Print(d)
		other.Print(", ")
		register.Print(s)
		other.Print(", ")
		register.Print(o)
	}

	rrn := func(name string, d cpu.Register, s cpu.Register, n int) {
		mnemonic.Print("  " + name)
		register.Print(d)
		other.Print(", ")
		register.Print(s)
		other.Print(", ")
		imm.Print(n)
	}

	rr := func(name string, d any, s any) {
		mnemonic.Print("  " + name)
		register.Print(d)
		other.Print(", ")
		register.Print(s)
	}

	dn := func(name string, d cpu.Register, n int) {
		mnemonic.Print("  " + name)
		register.Print(d)
		other.Print(", ")
		imm.Print(n)
	}

	regs := func(name string, registers []cpu.Register) {
		mnemonic.Print("  " + name)

		for i, reg := range registers {
			if i != 0 {
				other.Print(", ")
			}

			register.Print(reg)
		}
	}

	printValue := func(value any) {
		switch value := value.(type) {
		case cpu.Register:
			register.Print(value)
		case int, int32:
			imm.Print(value)
		}
	}

	load := func(length byte, d cpu.Register, base cpu.Register, index any, scale bool) {
		mnemonic.Printf("  load %db ", length)
		register.Print(d)
		other.Print(", [")
		register.Print(base)
		other.Print(" + ")
		printValue(index)

		if scale {
			other.Print(" * ")
			imm.Print(length)
		}

		other.Print("]")
	}

	store := func(length byte, base cpu.Register, index any, source any, scale bool) {
		mnemonic.Printf("  store %db ", length)
		other.Print("[")
		register.Print(base)
		other.Print(" + ")
		printValue(index)

		if scale {
			other.Print(" * ")
			imm.Print(length)
		}

		other.Print("], ")
		printValue(source)
	}

	condition := func(c asm.Condition) string {
		switch c {
		case asm.Equal:
			return "if == "
		case asm.NotEqual:
			return "if != "
		case asm.Greater:
			return "if.s > "
		case asm.GreaterEqual:
			return "if.s >= "
		case asm.Less:
			return "if.s < "
		case asm.LessEqual:
			return "if.s <= "
		case asm.UnsignedGreater:
			return "if.u > "
		case asm.UnsignedGreaterEqual:
			return "if.u >= "
		case asm.UnsignedLess:
			return "if.u < "
		case asm.UnsignedLessEqual:
			return "if.u <= "
		default:
			return ""
		}
	}

	for _, instr := range f.Assembler.Instructions {
		switch instr := instr.(type) {
		case *asm.Add:
			rrr("add ", instr.Destination, instr.Source, instr.Operand)
		case *asm.AddNumber:
			rrn("add ", instr.Destination, instr.Source, instr.Number)
		case *asm.And:
			rrr("and ", instr.Destination, instr.Source, instr.Operand)
		case *asm.AndNumber:
			rrn("and ", instr.Destination, instr.Source, instr.Number)
		case *asm.Call:
			mnemonic.Print("  call ")
			label.Print(ssa.CleanLabel(instr.Label))
		case *asm.CallExtern:
			mnemonic.Print("  call extern ")
			label.Print(instr.Library + "." + instr.Function)
		case *asm.Compare:
			rr("compare ", instr.Destination, instr.Source)
		case *asm.CompareAndSwap:
			rrr("cas ", instr.OldValue, instr.NewValue, instr.Address)
		case *asm.CompareNumber:
			dn("compare ", instr.Destination, instr.Number)
		case *asm.ConditionalSet:
			if suffix := condition(instr.Condition); suffix != "" {
				mnemonic.Print("  set " + suffix)
			} else {
				ansi.Red.Print("  set: unknown condition: " + fmt.Sprint(instr.Condition) + " ")
			}

			register.Print(instr.Destination)
		case *asm.Divide:
			rrr("div.u ", instr.Destination, instr.Source, instr.Operand)
		case *asm.DivideSigned:
			rrr("div.s ", instr.Destination, instr.Source, instr.Operand)
		case *asm.Jump:
			switch instr.Condition {
			case asm.None:
				mnemonic.Print("  jump ")
			default:
				suffix := condition(instr.Condition)

				if suffix != "" {
					mnemonic.Print("  jump " + suffix)
				} else {
					ansi.Red.Print("  jump: unknown condition: " + fmt.Sprint(instr.Condition) + " ")
				}
			}

			label.Print(ssa.CleanLabel(instr.Label))
		case *asm.Label:
			if instr.Name == f.FullName {
				function.Printf("%s:", instr.Name)
			} else {
				label.Printf("\n%s:", ssa.CleanLabel(instr.Name))
			}
		case *asm.Load:
			load(instr.Length, instr.Destination, instr.Base, instr.Index, instr.Scale)
		case *asm.LoadFixedOffset:
			load(instr.Length, instr.Destination, instr.Base, instr.Index, instr.Scale)
		case *asm.Modulo:
			rrr("mod.u ", instr.Destination, instr.Source, instr.Operand)
		case *asm.ModuloSigned:
			rrr("mod.s ", instr.Destination, instr.Source, instr.Operand)
		case *asm.Move:
			rr("move ", instr.Destination, instr.Source)
		case *asm.MoveLabel:
			mnemonic.Print("  address ")
			register.Print(instr.Destination)
			other.Print(", ")
			label.Print(ssa.CleanLabel(instr.Label))
		case *asm.MoveNumber:
			dn("move ", instr.Destination, instr.Number)
		case *asm.Multiply:
			rrr("mul ", instr.Destination, instr.Source, instr.Operand)
		case *asm.Negate:
			rr("neg ", instr.Destination, instr.Source)
		case *asm.Or:
			rrr("or ", instr.Destination, instr.Source, instr.Operand)
		case *asm.OrNumber:
			rrn("or ", instr.Destination, instr.Source, instr.Number)
		case *asm.Pop:
			regs("pop ", instr.Registers)
		case *asm.Push:
			regs("push ", instr.Registers)
		case *asm.ReadSystemRegister:
			rr("move ", instr.Destination, instr.SystemRegister)
		case *asm.Return:
			mnemonic.Print("  return")
		case *asm.ShiftLeft:
			rrr("shift << ", instr.Destination, instr.Source, instr.Operand)
		case *asm.ShiftLeftNumber:
			rrn("shift << ", instr.Destination, instr.Source, instr.Number)
		case *asm.ShiftRight:
			rrr("shift.u >> ", instr.Destination, instr.Source, instr.Operand)
		case *asm.ShiftRightNumber:
			rrn("shift.u >> ", instr.Destination, instr.Source, instr.Number)
		case *asm.ShiftRightSigned:
			rrr("shift.s >> ", instr.Destination, instr.Source, instr.Operand)
		case *asm.ShiftRightSignedNumber:
			rrn("shift.s >> ", instr.Destination, instr.Source, instr.Number)
		case *asm.Store:
			store(instr.Length, instr.Base, instr.Index, instr.Source, instr.Scale)
		case *asm.StoreFixedOffset:
			store(instr.Length, instr.Base, instr.Index, instr.Source, instr.Scale)
		case *asm.StoreFixedOffsetNumber:
			store(instr.Length, instr.Base, instr.Index, instr.Number, instr.Scale)
		case *asm.StoreNumber:
			store(instr.Length, instr.Base, instr.Index, instr.Number, instr.Scale)
		case *asm.Subtract:
			rrr("sub ", instr.Destination, instr.Source, instr.Operand)
		case *asm.SubtractNumber:
			rrn("sub ", instr.Destination, instr.Source, instr.Number)
		case *asm.StackFrameStart:
			mnemonic.Print("  frame start ")

			if instr.FramePointer {
				other.Print("fp ")
			}

			if instr.ExternCalls {
				other.Print("extern ")
			}
		case *asm.StackFrameEnd:
			mnemonic.Print("  frame end ")

			if instr.FramePointer {
				other.Print("fp ")
			}
		case *asm.Syscall:
			mnemonic.Print("  syscall")
		case *asm.WriteSystemRegister:
			rr("move ", instr.SystemRegister, instr.Source)
		case *asm.Xor:
			rrr("xor ", instr.Destination, instr.Source, instr.Operand)
		case *asm.XorNumber:
			rrn("xor ", instr.Destination, instr.Source, instr.Number)
		default:
			ansi.Red.Print("  unknown: " + reflect.TypeOf(instr).String())
		}

		fmt.Println()
	}

	fmt.Println()
}