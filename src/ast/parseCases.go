package ast

import (
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/fs"
	"git.urbach.dev/cli/q/src/token"
)

// parseCases generates the cases inside a switch statement.
func parseCases(tokens token.List, file *fs.File) ([]Case, error) {
	cases := make([]Case, 0, 4)

	for caseTokens := range tokens.Instructions {
		blockStart, _, body, err := block(caseTokens, file)

		if err != nil {
			return nil, err
		}

		conditionTokens := caseTokens[:blockStart]
		var condition *expression.Expression

		if len(conditionTokens) == 1 && conditionTokens[0].Kind == token.Identifier && conditionTokens[0].StringFrom(file.Bytes) == "_" {
			condition = nil
		} else {
			condition = expression.Parse(conditionTokens)
		}

		cases = append(cases, Case{
			Condition: condition,
			Body:      body,
		})
	}

	return cases, nil
}