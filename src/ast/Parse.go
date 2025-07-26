package ast

import (
	"git.urbach.dev/cli/q/src/fs"
	"git.urbach.dev/cli/q/src/token"
)

// Parse generates an AST from a list of tokens.
func Parse(tokens token.List, file *fs.File) (AST, error) {
	nodes := make(AST, 0, len(tokens)/64)

	for tokens := range tokens.Instructions {
		node, err := parseInstruction(tokens, file, nodes)

		if node != nil {
			nodes = append(nodes, node)
		}

		if err != nil {
			return nil, err
		}
	}

	return nodes, nil
}