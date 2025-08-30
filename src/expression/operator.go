package expression

import (
	"math"

	"git.urbach.dev/cli/q/src/token"
)

// operator represents an operator for mathematical expressions.
type operator struct {
	Precedence int8
	Operands   int8
}

// Operators defines the Operators used in the language.
// The number corresponds to the operator priority and can not be zero.
var Operators = [64]operator{
	token.Struct: {13, 1},
	token.Dot:    {13, 2},
	token.Call:   {12, 1},
	token.Array:  {12, 2},
	token.Negate: {11, 1},
	token.Not:    {11, 1},
	token.Mul:    {10, 2},
	token.Div:    {10, 2},
	token.Mod:    {10, 2},
	token.Add:    {9, 2},
	token.Sub:    {9, 2},
	token.Shr:    {8, 2},
	token.Shl:    {8, 2},
	token.And:    {7, 2},
	token.Xor:    {6, 2},
	token.Or:     {5, 2},

	token.Greater:      {4, 2},
	token.Less:         {4, 2},
	token.GreaterEqual: {4, 2},
	token.LessEqual:    {4, 2},
	token.Equal:        {3, 2},
	token.NotEqual:     {3, 2},
	token.LogicalAnd:   {2, 2},
	token.LogicalOr:    {1, 2},

	token.Range:     {0, 2},
	token.Separator: {0, 2},

	token.Assign:      {math.MinInt8, 2},
	token.Define:      {math.MinInt8, 2},
	token.AddAssign:   {math.MinInt8, 2},
	token.SubAssign:   {math.MinInt8, 2},
	token.MulAssign:   {math.MinInt8, 2},
	token.DivAssign:   {math.MinInt8, 2},
	token.ModAssign:   {math.MinInt8, 2},
	token.AndAssign:   {math.MinInt8, 2},
	token.OrAssign:    {math.MinInt8, 2},
	token.XorAssign:   {math.MinInt8, 2},
	token.ShrAssign:   {math.MinInt8, 2},
	token.ShlAssign:   {math.MinInt8, 2},
	token.FieldAssign: {math.MinInt8, 2},
}

func numOperands(symbol token.Kind) int {
	return int(Operators[symbol].Operands)
}

func precedence(symbol token.Kind) int8 {
	return Operators[symbol].Precedence
}