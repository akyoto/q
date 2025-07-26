package token

import (
	"unsafe"
)

// List is a slice of tokens.
type List []Token

// IndexKind returns the position of a token kind within a token list.
func (list List) IndexKind(kind Kind) int {
	for i, token := range list {
		if token.Kind == kind {
			return i
		}
	}

	return -1
}

// LastIndexKind returns the position of the last token kind within a token list.
func (list List) LastIndexKind(kind Kind) int {
	for i := len(list) - 1; i >= 0; i-- {
		if list[i].Kind == kind {
			return i
		}
	}

	return -1
}

// Split calls the callback function on each set of tokens in a comma separated list.
func (list List) Split(yield func(List) bool) {
	if len(list) == 0 {
		return
	}

	start := 0
	groupLevel := 0

	for i, t := range list {
		switch t.Kind {
		case GroupStart, ArrayStart, BlockStart:
			groupLevel++

		case GroupEnd, ArrayEnd, BlockEnd:
			groupLevel--

		case Separator:
			if groupLevel > 0 {
				continue
			}

			parameter := list[start:i]

			if !yield(parameter) {
				return
			}

			start = i + 1
		}
	}

	yield(list[start:])
}

// String returns the concatenated token strings.
func (list List) String(source []byte) string {
	start := list[0].Position
	end := list[len(list)-1].End()
	return unsafe.String(unsafe.SliceData(source[start:end]), end-start)
}