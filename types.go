package main

import "github.com/akyoto/q/token"

// type Function interface {
// 	Name() string
// 	Parameters() []Expression
// }

type FunctionCall struct {
	FunctionName   string
	Parameters     []Expression
	ParameterStart int
}

type Expression []token.Token
