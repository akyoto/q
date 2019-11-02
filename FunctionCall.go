package main

type FunctionCall struct {
	FunctionName   string
	Parameters     []Expression
	ParameterStart int
}

type Expression []Token
