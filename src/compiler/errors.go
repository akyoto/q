package compiler

import "git.urbach.dev/cli/q/src/errors"

var (
	MissingInitFunction = errors.String("Missing init function")
	MissingMainFunction = errors.String("Missing main function")
)