package fs

import "git.urbach.dev/cli/q/src/token"

// Import contains data about an import statement in a file.
type Import struct {
	Package  string
	Position token.Position
	Used     bool
}