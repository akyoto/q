package fs

import "git.urbach.dev/cli/q/src/token"

// File represents a single source file.
type File struct {
	Path    string
	Package string
	Imports map[string]*Import
	Bytes   []byte
	Tokens  token.List
}