package expression

import (
	"git.urbach.dev/cli/q/src/token"
)

// operator represents an operator for mathematical expressions.
type operator struct {
	Precedence int8
	Operands   int8
}

// Operators defines the Operators used in the language.
// The number corresponds to the operator priority and can not be zero.
var Operators = [256]operator{
	token.Dot:    {8, 2},
	token.Call:   {8, 1},
	token.Array:  {8, 2},
	token.Struct: {8, 1},

	token.Negate: {7, 1},
	token.Not:    {7, 1},

	token.Mul: {6, 2},
	token.Div: {6, 2},
	token.Mod: {6, 2},

	token.Add:  {5, 2},
	token.Sub:  {5, 2},
	token.And:  {5, 2},
	token.Or:   {5, 2},
	token.Xor:  {5, 2},
	token.Shl:  {5, 2},
	token.Shr:  {5, 2},
	token.Cast: {5, 2},

	token.Greater:      {4, 2},
	token.Less:         {4, 2},
	token.GreaterEqual: {4, 2},
	token.LessEqual:    {4, 2},
	token.Equal:        {4, 2},
	token.NotEqual:     {4, 2},

	token.LogicalAnd: {3, 2},
	token.LogicalOr:  {3, 2},

	token.Range:     {2, 2},
	token.Separator: {2, 2},

	token.Assign:      {1, 2},
	token.Define:      {1, 2},
	token.AddAssign:   {1, 2},
	token.SubAssign:   {1, 2},
	token.MulAssign:   {1, 2},
	token.DivAssign:   {1, 2},
	token.ModAssign:   {1, 2},
	token.AndAssign:   {1, 2},
	token.OrAssign:    {1, 2},
	token.XorAssign:   {1, 2},
	token.ShrAssign:   {1, 2},
	token.ShlAssign:   {1, 2},
	token.FieldAssign: {1, 2},
}

func numOperands(symbol token.Kind) int {
	return int(Operators[symbol].Operands)
}

func precedence(symbol token.Kind) int8 {
	return Operators[symbol].Precedence
}