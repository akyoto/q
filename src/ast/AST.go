package ast

import "git.urbach.dev/cli/q/src/expression"

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
	For struct {
		Head *expression.Expression
		Body AST
	}
	If struct {
		Condition *expression.Expression
		Body      AST
		Else      AST
	}
	Loop struct {
		Body AST
	}
	Return struct {
		Values []*expression.Expression
	}
	Switch struct {
		Cases []Case
	}
)