package ast

import (
	"fmt"

	"git.urbach.dev/cli/q/src/fs"
	"git.urbach.dev/cli/q/src/token"
)

// parseKeyword generates a keyword node from an instruction.
func parseKeyword(tokens token.List, file *fs.File, nodes AST) (Node, error) {
	switch tokens[0].Kind {
	case token.Assert:
		return parseAssert(tokens, file)
	case token.Else:
		return parseElse(tokens, file, nodes)
	case token.Go:
		return parseGo(tokens, file)
	case token.If:
		return parseIf(tokens, file)
	case token.Loop:
		return parseLoop(tokens, file)
	case token.Return:
		return parseReturn(tokens, file)
	case token.Switch:
		return parseSwitch(tokens, file)
	default:
		panic(fmt.Sprintf("keyword not implemented: %s", tokens[0].String(file.Bytes)))
	}
}