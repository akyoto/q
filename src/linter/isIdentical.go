package linter

import (
	"git.urbach.dev/cli/q/src/ssa"
)

// isIdentical returns true if the values are identical or copies of identical values.
func isIdentical(left ssa.Value, right ssa.Value) bool {
	if left == right {
		return true
	}

	leftCopy, leftIsCopy := left.(*ssa.Copy)

	if leftIsCopy && leftCopy.Value == right {
		return true
	}

	rightCopy, rightIsCopy := right.(*ssa.Copy)

	if rightIsCopy && rightCopy.Value == left {
		return true
	}

	if leftIsCopy && rightIsCopy && leftCopy.Value == rightCopy.Value {
		return true
	}

	return false
}