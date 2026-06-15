package optimizer

import (
	"slices"

	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/token"
)

// Complexity returns the number of registers needed to evaluate the expression.
func Complexity(expr *expression.Expression) int {
	if expr.IsLeaf() {
		return 1
	}

	switch expr.Token.Kind {
	case token.Call:
		return complexityKeepAlive(expr.Children[1:])

	case token.Dot:
		return Complexity(expr.Children[0])

	case token.Array:
		return complexityKeepAlive(expr.Children)

	case token.Struct:
		if expr.Parent != nil && expr.Parent.Token.Kind == token.New {
			return complexityNoKeepAlive(expr.Children[1:])
		}

		return complexityKeepAlive(expr.Children[1:])

	case token.FieldAssign:
		return Complexity(expr.Children[1])

	case token.Cast:
		return Complexity(expr.Children[0])
	}

	if len(expr.Children) == 1 {
		return Complexity(expr.Children[0])
	}

	return complexityBinary(expr)
}

// complexityBinary returns the number of registers needed to evaluate the binary operation.
func complexityBinary(expr *expression.Expression) int {
	left := Complexity(expr.Children[0])
	right := Complexity(expr.Children[1])

	if left == right {
		return left + 1
	}

	return max(left, right)
}

// complexityKeepAlive returns the number of registers needed
// when all expressions must be kept alive at the same time.
func complexityKeepAlive(args []*expression.Expression) int {
	if len(args) == 0 {
		return 1
	}

	weights := make([]int, len(args))

	for i, arg := range args {
		weights[i] = Complexity(arg)
	}

	slices.SortFunc(weights, func(a int, b int) int {
		return b - a
	})

	maxRegs := 0

	for i, weight := range weights {
		regs := weight + i

		if regs > maxRegs {
			maxRegs = regs
		}
	}

	return maxRegs
}

// complexityNoKeepAlive returns the number of registers needed
// when each expression is no longer needed after its evaluation.
func complexityNoKeepAlive(args []*expression.Expression) int {
	maxRegs := 0

	for _, arg := range args {
		regs := Complexity(arg)

		if regs > maxRegs {
			maxRegs = regs
		}
	}

	return maxRegs
}