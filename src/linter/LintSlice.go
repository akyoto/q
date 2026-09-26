package linter

import (
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/fs"
	"git.urbach.dev/cli/q/src/token"
)

// LintSlice checks for simplifications of the lower and upper range of slices.
func LintSlice(expr *expression.Expression, index *expression.Expression, file *fs.File) error {
	if len(index.Children) == 0 {
		return nil
	}

	if isZero(index.Children[0], file) {
		return simplifySlice(expr, index.Children[0], file)
	}

	if len(index.Children) == 2 && isLengthOf(expr.Children[0], index.Children[1], file) {
		return simplifySlice(expr, index.Children[1], file)
	}

	return nil
}

// isLengthOf returns true if the dot expression is the length of the base.
func isLengthOf(base *expression.Expression, dot *expression.Expression, file *fs.File) bool {
	if dot.Token.Kind != token.Dot || len(dot.Children) != 2 {
		return false
	}

	return dot.Children[0].SourceString(file.Bytes) == base.SourceString(file.Bytes) && dot.Children[1].SourceString(file.Bytes) == "len"
}

// isZero returns true if the expression is the literal 0.
func isZero(expr *expression.Expression, file *fs.File) bool {
	return expr.Token.Kind == token.Number && expr.SourceString(file.Bytes) == "0"
}

// simplifySlice returns an error that suggests removing an operand from the slice range.
func simplifySlice(expr *expression.Expression, limit *expression.Expression, file *fs.File) error {
	text := expr.SourceString(file.Bytes)
	offset := expr.Source().Start()
	text = text[:limit.Source().Start()-offset] + text[limit.Source().End()-offset:]
	return errors.New(&Simplify{To: text}, file, expr.Source())
}