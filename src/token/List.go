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
func (list List) Split(yield func(Position, List) bool) {
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

			if !yield(list[start].Position, list[start:i]) {
				return
			}

			start = i + 1
		}
	}

	if start < len(list) {
		yield(list[start].Position, list[start:])
		return
	}

	yield(list[len(list)-1].End(), list[start:])
}

func (list List) Start() Position {
	return list[0].Position
}

func (list List) End() Position {
	return list[len(list)-1].End()
}

// StringFrom returns the source string from the first to the last token.
func (list List) StringFrom(source []byte) string {
	if len(list) == 0 {
		return ""
	}

	start := list.Start()
	end := list.End()
	return unsafe.String(unsafe.SliceData(source[start:end]), end-start)
}