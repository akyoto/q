package ast

import (
	"git.urbach.dev/cli/q/src/errors"
)

var (
	EmptySwitch          = errors.String("Empty switch")
	ExpectedBlock        = errors.String("Expected '{'")
	ExpectedIfBeforeElse = errors.String("Expected an 'if' block before 'else'")
	InvalidInstruction   = errors.String("Invalid instruction")
	MissingBlockStart    = errors.String("Missing '{'")
	MissingBlockEnd      = errors.String("Missing '}'")
	MissingExpression    = errors.String("Missing expression")
	MissingOperand       = errors.String("Missing operand")
	NoElseIf             = errors.String("Use 'switch' instead of 'else if'")
)