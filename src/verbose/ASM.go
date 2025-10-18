package verbose

import (
	_ "embed"
	"fmt"
	"reflect"

	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/core"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/go/color/ansi"
)

//go:embed ASM.txt
var HeaderASM string

// ASM shows the assembly code.
func ASM(root *core.Function) {
	root.EachDependency(make(map[*core.Function]bool), func(f *core.Function) {
		if filter(f.FullName, f.Env.Build.Filter) {
			return
		}

		printAssembly(f)
	})
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

	for _, instr := range f.Assembler.Instructions {
		switch instr := instr.(type) {
		case *asm.Add:
			mnemonic.Print("  add ")
			register.Print(instr.Destination)
			other.Print(", ")
			register.Print(instr.Source)
			other.Print(", ")
			register.Print(instr.Operand)
		case *asm.AddNumber:
			mnemonic.Print("  add ")
			register.Print(instr.Destination)
			other.Print(", ")
			register.Print(instr.Source)
			other.Print(", ")
			imm.Print(instr.Number)
		case *asm.And:
			mnemonic.Print("  and ")
			register.Print(instr.Destination)
			other.Print(", ")
			register.Print(instr.Source)
			other.Print(", ")
			register.Print(instr.Operand)
		case *asm.AndNumber:
			mnemonic.Print("  and ")
			register.Print(instr.Destination)
			other.Print(", ")
			register.Print(instr.Source)
			other.Print(", ")
			imm.Print(instr.Number)
		case *asm.Call:
			mnemonic.Print("  call ")
			label.Print(ssa.CleanLabel(instr.Label))
		case *asm.CallExtern:
			mnemonic.Print("  call extern ")
			label.Print(instr.Library + "." + instr.Function)
		case *asm.Compare:
			mnemonic.Print("  compare ")
			register.Print(instr.Destination)
			other.Print(", ")
			register.Print(instr.Source)
		case *asm.CompareNumber:
			mnemonic.Print("  compare ")
			register.Print(instr.Destination)
			other.Print(", ")
			imm.Print(instr.Number)
		case *asm.Divide:
			mnemonic.Print("  div.u ")
			register.Print(instr.Destination)
			other.Print(", ")
			register.Print(instr.Source)
			other.Print(", ")
			register.Print(instr.Operand)
		case *asm.DivideSigned:
			mnemonic.Print("  div.s ")
			register.Print(instr.Destination)
			other.Print(", ")
			register.Print(instr.Source)
			other.Print(", ")
			register.Print(instr.Operand)
		case *asm.Jump:
			switch instr.Condition {
			case token.Equal:
				mnemonic.Print("  jump if == ")
			case token.Greater:
				mnemonic.Print("  jump if > ")
			case token.GreaterEqual:
				mnemonic.Print("  jump if >= ")
			case token.Less:
				mnemonic.Print("  jump if < ")
			case token.LessEqual:
				mnemonic.Print("  jump if <= ")
			case token.NotEqual:
				mnemonic.Print("  jump if != ")
			case token.Invalid:
				mnemonic.Print("  jump ")
			}

			label.Print(ssa.CleanLabel(instr.Label))
		case *asm.Label:
			if instr.Name == f.FullName {
				function.Printf("%s:", instr.Name)
			} else {
				label.Printf("\n%s:", ssa.CleanLabel(instr.Name))
			}
		case *asm.Load:
			mnemonic.Printf("  load %db ", instr.Length)
			register.Print(instr.Destination)
			other.Print(", [")
			register.Print(instr.Base)
			other.Print(" + ")
			register.Print(instr.Index)
			if instr.Scale {
				other.Print(" * ")
				imm.Print(instr.Length)
			}
			other.Print("]")
		case *asm.LoadFixedOffset:
			mnemonic.Printf("  load %db ", instr.Length)
			register.Print(instr.Destination)
			other.Print(", [")
			register.Print(instr.Base)
			other.Print(" + ")
			imm.Print(instr.Index)
			if instr.Scale {
				other.Print(" * ")
				imm.Print(instr.Length)
			}
			other.Print("]")
		case *asm.Modulo:
			mnemonic.Print("  mod.u ")
			register.Print(instr.Destination)
			other.Print(", ")
			register.Print(instr.Source)
			other.Print(", ")
			register.Print(instr.Operand)
		case *asm.ModuloSigned:
			mnemonic.Print("  mod.s ")
			register.Print(instr.Destination)
			other.Print(", ")
			register.Print(instr.Source)
			other.Print(", ")
			register.Print(instr.Operand)
		case *asm.MoveLabel:
			mnemonic.Print("  address ")
			register.Print(instr.Destination)
			other.Print(", ")
			label.Print(ssa.CleanLabel(instr.Label))
		case *asm.Move:
			mnemonic.Print("  move ")
			register.Print(instr.Destination)
			other.Print(", ")
			register.Print(instr.Source)
		case *asm.MoveNumber:
			mnemonic.Print("  move ")
			register.Print(instr.Destination)
			other.Print(", ")
			imm.Print(instr.Number)
		case *asm.Multiply:
			mnemonic.Print("  mul ")
			register.Print(instr.Destination)
			other.Print(", ")
			register.Print(instr.Source)
			other.Print(", ")
			register.Print(instr.Operand)
		case *asm.Negate:
			mnemonic.Print("  neg ")
			register.Print(instr.Destination)
			other.Print(", ")
			register.Print(instr.Source)
		case *asm.Or:
			mnemonic.Print("  or ")
			register.Print(instr.Destination)
			other.Print(", ")
			register.Print(instr.Source)
			other.Print(", ")
			register.Print(instr.Operand)
		case *asm.OrNumber:
			mnemonic.Print("  or ")
			register.Print(instr.Destination)
			other.Print(", ")
			register.Print(instr.Source)
			other.Print(", ")
			imm.Print(instr.Number)
		case *asm.Pop:
			mnemonic.Print("  pop ")
			for i, reg := range instr.Registers {
				if i != 0 {
					other.Print(", ")
				}
				register.Print(reg)
			}
		case *asm.Push:
			mnemonic.Print("  push ")
			for i, reg := range instr.Registers {
				if i != 0 {
					other.Print(", ")
				}
				register.Print(reg)
			}
		case *asm.Return:
			mnemonic.Print("  return")
		case *asm.ShiftLeft:
			mnemonic.Print("  shift << ")
			register.Print(instr.Destination)
			other.Print(", ")
			register.Print(instr.Source)
			other.Print(", ")
			register.Print(instr.Operand)
		case *asm.ShiftLeftNumber:
			mnemonic.Print("  shift << ")
			register.Print(instr.Destination)
			other.Print(", ")
			register.Print(instr.Source)
			other.Print(", ")
			imm.Print(instr.Number)
		case *asm.ShiftRight:
			mnemonic.Print("  shift.u >> ")
			register.Print(instr.Destination)
			other.Print(", ")
			register.Print(instr.Source)
			other.Print(", ")
			register.Print(instr.Operand)
		case *asm.ShiftRightNumber:
			mnemonic.Print("  shift.u >> ")
			register.Print(instr.Destination)
			other.Print(", ")
			register.Print(instr.Source)
			other.Print(", ")
			imm.Print(instr.Number)
		case *asm.ShiftRightSigned:
			mnemonic.Print("  shift.s >> ")
			register.Print(instr.Destination)
			other.Print(", ")
			register.Print(instr.Source)
			other.Print(", ")
			register.Print(instr.Operand)
		case *asm.ShiftRightSignedNumber:
			mnemonic.Print("  shift.s >> ")
			register.Print(instr.Destination)
			other.Print(", ")
			register.Print(instr.Source)
			other.Print(", ")
			imm.Print(instr.Number)
		case *asm.Store:
			mnemonic.Printf("  store %db ", instr.Length)
			other.Print("[")
			register.Print(instr.Base)
			other.Print(" + ")
			register.Print(instr.Index)
			if instr.Scale {
				other.Print(" * ")
				imm.Print(instr.Length)
			}
			other.Print("], ")
			register.Print(instr.Source)
		case *asm.StoreFixedOffset:
			mnemonic.Printf("  store %db ", instr.Length)
			other.Print("[")
			register.Print(instr.Base)
			other.Print(" + ")
			imm.Print(instr.Index)
			if instr.Scale {
				other.Print(" * ")
				imm.Print(instr.Length)
			}
			other.Print("], ")
			register.Print(instr.Source)
		case *asm.StoreFixedOffsetNumber:
			mnemonic.Printf("  store %db ", instr.Length)
			other.Print("[")
			register.Print(instr.Base)
			other.Print(" + ")
			imm.Print(instr.Index)
			if instr.Scale {
				other.Print(" * ")
				imm.Print(instr.Length)
			}
			other.Print("], ")
			imm.Print(instr.Number)
		case *asm.StoreNumber:
			mnemonic.Printf("  store %db ", instr.Length)
			other.Print("[")
			register.Print(instr.Base)
			other.Print(" + ")
			register.Print(instr.Index)
			if instr.Scale {
				other.Print(" * ")
				imm.Print(instr.Length)
			}
			other.Print("], ")
			imm.Print(instr.Number)
		case *asm.Subtract:
			mnemonic.Print("  sub ")
			register.Print(instr.Destination)
			other.Print(", ")
			register.Print(instr.Source)
			other.Print(", ")
			register.Print(instr.Operand)
		case *asm.SubtractNumber:
			mnemonic.Print("  sub ")
			register.Print(instr.Destination)
			other.Print(", ")
			register.Print(instr.Source)
			other.Print(", ")
			imm.Print(instr.Number)
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
		case *asm.Xor:
			mnemonic.Print("  xor ")
			register.Print(instr.Destination)
			other.Print(", ")
			register.Print(instr.Source)
			other.Print(", ")
			register.Print(instr.Operand)
		case *asm.XorNumber:
			mnemonic.Print("  xor ")
			register.Print(instr.Destination)
			other.Print(", ")
			register.Print(instr.Source)
			other.Print(", ")
			imm.Print(instr.Number)
		default:
			ansi.Red.Print("  unknown: " + reflect.TypeOf(instr).String())
		}

		fmt.Println()
	}

	fmt.Println()
}