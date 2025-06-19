package token

// Instructions yields on each AST node.
func (list List) Instructions(yield func(List) bool) {
	start := 0
	groupLevel := 0
	blockLevel := 0

	for i, t := range list {
		switch t.Kind {
		case NewLine:
			if start == i {
				start = i + 1
				continue
			}

			if groupLevel > 0 || blockLevel > 0 {
				continue
			}

			if !yield(list[start:i]) {
				return
			}

			start = i + 1

		case GroupStart:
			groupLevel++

		case GroupEnd:
			groupLevel--

		case BlockStart:
			blockLevel++

		case BlockEnd:
			blockLevel--

			if groupLevel > 0 || blockLevel > 0 {
				continue
			}

			if !yield(list[start : i+1]) {
				return
			}

			start = i + 1

		case EOF:
			if start < i {
				yield(list[start:i])
			}

			return
		}
	}

	if start < len(list) {
		yield(list[start:])
	}
}