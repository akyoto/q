package ast

import (
	"git.urbach.dev/cli/q/src/errors"
)

var (
	EmptySwitch          = errors.String("Empty switch")
	ExpectedIfBeforeElse = errors.String("Expected an 'if' block before 'else'")
	InvalidInstruction   = errors.String("Invalid instruction")
	MissingBlockStart    = errors.String("Missing '{'")
	MissingBlockEnd      = errors.String("Missing '}'")
	MissingExpression    = errors.String("Missing expression")
	MissingOperand       = errors.String("Missing operand")
)