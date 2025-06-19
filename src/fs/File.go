package fs

// File represents a single source file.
type File struct {
	Path    string
	Package string
	Bytes   []byte
}