package core

import (
	"strconv"
	"strings"
	"unicode/utf8"

	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/fs"
	"git.urbach.dev/cli/q/src/token"
)

// toNumber tries to convert the token into a numeric value.
func (f *Function) toNumber(t token.Token) (int, error) {
	return toNumber(t, f.File)
}

// toNumber tries to convert the token into a numeric value.
func toNumber(t token.Token, file *fs.File) (int, error) {
	switch t.Kind {
	case token.Number:
		var (
			digits = t.String(file.Bytes)
			number int64
			err    error
		)

		switch {
		case strings.HasPrefix(digits, "0x"):
			number, err = strconv.ParseInt(digits[2:], 16, 64)
		case strings.HasPrefix(digits, "0o"):
			number, err = strconv.ParseInt(digits[2:], 8, 64)
		case strings.HasPrefix(digits, "0b"):
			number, err = strconv.ParseInt(digits[2:], 2, 64)
		default:
			number, err = strconv.ParseInt(digits, 10, 64)
		}

		if err != nil {
			return 0, errors.New(InvalidNumber, file, t.Position)
		}

		return int(number), nil

	case token.Rune:
		r := t.Bytes(file.Bytes)
		r = unescape(r)

		if len(r) == 0 {
			return 0, errors.New(InvalidRune, file, t.Position+1)
		}

		number, size := utf8.DecodeRune(r)

		if len(r) > size {
			return 0, errors.New(InvalidRune, file, t.Position+1)
		}

		return int(number), nil
	}

	return 0, errors.New(InvalidNumber, file, t.Position)
}