package verbose

import (
	_ "embed"
	"fmt"
	"strings"

	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/core"
	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/go/color/ansi"
)

//go:embed ASM.txt
var HeaderASM string

// ASM shows the assembly code.
func ASM(root *core.Function) {
	root.EachDependency(make(map[*core.Function]bool), func(f *core.Function) {
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
			label.Print(instr.Label)
		case *asm.CallExtern:
			mnemonic.Print("  call extern ")
			label.Print(instr.Library + "." + instr.Function)
		case *asm.Compare:
			mnemonic.Print("  compare ")
			register.Print(instr.Destination)
			other.Print(", ")
			register.Print(instr.Source)
		case *asm.Divide:
			mnemonic.Print("  div ")
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

			label.Print(strings.TrimPrefix(instr.Label, f.FullName))
		case *asm.Label:
			if instr.Name == f.FullName {
				function.Printf("%s:", instr.Name)
			} else {
				label.Printf("\n%s:", strings.TrimPrefix(instr.Name, f.FullName))
			}
		case *asm.Modulo:
		case *asm.MoveLabel:
			mnemonic.Print("  address ")
			register.Print(instr.Destination)
			other.Print(", ")
			label.Print(instr.Label)
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
		case *asm.Pop:
			mnemonic.Print("  pop ")
			register.Print(instr.Registers)
		case *asm.Push:
			mnemonic.Print("  push ")
			register.Print(instr.Registers)
		case *asm.Return:
			mnemonic.Print("  return")
		case *asm.ShiftLeft:
		case *asm.ShiftRightSigned:
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
		default:
			mnemonic.Print("  unknown")
		}

		fmt.Println()
	}

	fmt.Println()
}