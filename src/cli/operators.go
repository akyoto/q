package cli

import (
	"fmt"
	"slices"

	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/token"
)

// operators shows the entire list of operators grouped by precedence.
func operators() int {
	type operator struct {
		kind       token.Kind
		precedence int8
	}

	var ops []operator

	for i, op := range expression.Operators {
		if op.Operands == 0 {
			continue
		}

		ops = append(ops, operator{
			kind:       token.Kind(i),
			precedence: op.Precedence,
		})
	}

	slices.SortStableFunc(ops, func(a operator, b operator) int {
		return int(b.precedence) - int(a.precedence)
	})

	lastPrecedence := ops[0].precedence
	level := 0
	fmt.Printf("[%d] ", level)

	for _, op := range ops {
		var text string

		switch op.kind {
		case token.Call:
			text = "call()"
		case token.Array:
			text = "[index]"
		case token.Struct:
			text = "struct{}"
		default:
			text = op.kind.String()
		}

		if text == "" {
			panic(fmt.Sprintf("invalid operator string for %d", int(op.kind)))
		}

		if op.precedence != lastPrecedence {
			level++
			fmt.Printf("\n[%d] ", level)
		}

		lastPrecedence = op.precedence
		fmt.Printf("%s ", text)
	}

	fmt.Println()
	return success
}