package core

import "bytes"

// Unescape replaces the escape sequences in the contents of a string token with the respective characters.
func Unescape(data []byte) []byte {
	data = data[1 : len(data)-1]
	escape := bytes.IndexByte(data, '\\')

	if escape == -1 {
		return data
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
		}

		data = data[escape+2:]
		escape = bytes.IndexByte(data, '\\')

		if escape == -1 {
			return tmp
		}
	}
}