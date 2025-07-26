package core

import (
	"git.urbach.dev/cli/q/src/ast"
)

// compileAST compiles an abstract syntax tree.
func (f *Function) compileAST(tree ast.AST) error {
	for _, node := range tree {
		err := f.compileASTNode(node)

		if err != nil {
			return err
		}
	}

	return nil
}