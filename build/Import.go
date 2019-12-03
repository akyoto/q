package build

import "github.com/akyoto/q/build/token"

// Import represents an import statement in a file.
type Import struct {
	Path     string
	FullPath string
	Position token.Position
	Used     int32
}
