package core

import (
	"bytes"

	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/fs"
	"git.urbach.dev/cli/q/src/token"
)

// unescape replaces the escape sequences in the contents of a string token with the respective characters.
func unescape(t token.Token, file *fs.File) ([]byte, error) {
	data := t.Bytes(file.Bytes)
	data = data[1 : len(data)-1]
	escape := bytes.IndexByte(data, '\\')

	if escape == -1 {
		return data, nil
	}

	tmp := make([]byte, 0, len(data))

	for {
		tmp = append(tmp, data[:escape]...)

		switch data[escape+1] {
		case '0':
			tmp = append(tmp, '\000')
		case 't':
			tmp = append(tmp, '\t')
		case 'n':
			tmp = append(tmp, '\n')
		case 'r':
			tmp = append(tmp, '\r')
		case '"':
			tmp = append(tmp, '"')
		case '\'':
			tmp = append(tmp, '\'')
		case '\\':
			tmp = append(tmp, '\\')
		default:
			sourceStart := t.Start() + token.Position(escape) + 1
			sourceEnd := t.Start() + token.Position(escape) + 3
			source := token.NewSource(sourceStart, sourceEnd)
			return nil, errors.New(InvalidEscapeSequence, file, source)
		}

		data = data[escape+2:]
		escape = bytes.IndexByte(data, '\\')

		if escape == -1 {
			return append(tmp, data...), nil
		}
	}
}