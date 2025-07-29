package ast

import (
	"fmt"

	"git.urbach.dev/cli/q/src/errors"
)

var (
	EmptySwitch          = errors.String("Empty switch")
	ExpectedIfBeforeElse = errors.String("Expected an 'if' block before 'else'")
	MissingBlockStart    = errors.String("Missing '{'")
	MissingBlockEnd      = errors.String("Missing '}'")
	MissingExpression    = errors.String("Missing expression")
	MissingOperand       = errors.String("Missing operand")
)

// InvalidInstruction error is created when an instruction is not valid.
type InvalidInstruction struct {
	Instruction string
}

func (err *InvalidInstruction) Error() string {
	return fmt.Sprintf("Invalid instruction '%s'", err.Instruction)
}