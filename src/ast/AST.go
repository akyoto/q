package ast

import (
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/token"
)

type (
	Node any
	AST  []Node
)

type (
	Assert struct {
		Condition *expression.Expression
	}
	Assign struct {
		Expression *expression.Expression
	}
	Call struct {
		Expression *expression.Expression
	}
	Case struct {
		Condition *expression.Expression
		Body      AST
	}
	Define struct {
		Expression *expression.Expression
	}
	Go struct {
		Call *expression.Expression
	}
	If struct {
		Condition *expression.Expression
		Body      AST
		Else      AST
	}
	Loop struct {
		Head *expression.Expression
		Body AST
	}
	Return struct {
		Values []*expression.Expression
		Token  token.Token
	}
	Switch struct {
		Cases []Case
	}
)