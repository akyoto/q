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
	token.Struct: {7, 1},
	token.Dot:    {7, 2},
	token.Call:   {7, 1},
	token.Array:  {7, 2},
	token.Negate: {6, 1},
	token.Not:    {6, 1},
	token.Mul:    {5, 2},
	token.Div:    {5, 2},
	token.Mod:    {5, 2},
	token.Add:    {4, 2},
	token.Sub:    {4, 2},
	token.Shr:    {4, 2},
	token.Shl:    {4, 2},
	token.And:    {4, 2},
	token.Xor:    {4, 2},
	token.Or:     {4, 2},
	token.Cast:   {4, 2},

	token.Greater:      {3, 2},
	token.Less:         {3, 2},
	token.GreaterEqual: {3, 2},
	token.LessEqual:    {3, 2},
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