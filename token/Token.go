package token

// Token represents a single element in a source file.
// The characters that make up an identifier are grouped into a single token.
// This makes parsing easier and allows us to do better syntax checks.
type Token struct {
	Kind Kind
	Text string
}
