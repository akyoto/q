package fs

import "git.urbach.dev/cli/q/src/token"

// NewFile creates a new source file object.
func NewFile(path string, pkg string, source []byte) *File {
	return &File{
		Path:    path,
		Package: pkg,
		Bytes:   source,
		Tokens:  token.Tokenize(source),
		Imports: make(map[string]*Import, 4),
	}
}