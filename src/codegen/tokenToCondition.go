package codegen

import (
	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/token"
)

func tokenToCondition(cond token.Kind, unsigned bool) asm.Condition {
	switch cond {
	case token.Invalid:
		return asm.None
	case token.Equal:
		return asm.Equal
	case token.NotEqual:
		return asm.NotEqual
	case token.Greater:
		if unsigned {
			return asm.UnsignedGreater
		}

		return asm.Greater
	case token.GreaterEqual:
		if unsigned {
			return asm.UnsignedGreaterEqual
		}

		return asm.GreaterEqual
	case token.Less:
		if unsigned {
			return asm.UnsignedLess
		}

		return asm.Less
	case token.LessEqual:
		if unsigned {
			return asm.UnsignedLessEqual
		}

		return asm.LessEqual
	default:
		panic("unknown condition")
	}
}