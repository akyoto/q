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
func toNumber(t token.Token, file *fs.File) (int, error) {
	switch t.Kind {
	case token.Number:
		var (
			digits   = t.StringFrom(file.Bytes)
			signed   int64
			unsigned uint64
			err      error
		)

		switch {
		case strings.HasPrefix(digits, "0x"):
			signed, err = strconv.ParseInt(digits[2:], 16, 64)
		case strings.HasPrefix(digits, "0o"):
			signed, err = strconv.ParseInt(digits[2:], 8, 64)
		case strings.HasPrefix(digits, "0b"):
			signed, err = strconv.ParseInt(digits[2:], 2, 64)
		default:
			signed, err = strconv.ParseInt(digits, 10, 64)
		}

		if err != nil {
			switch {
			case strings.HasPrefix(digits, "0x"):
				unsigned, err = strconv.ParseUint(digits[2:], 16, 64)
			case strings.HasPrefix(digits, "0o"):
				unsigned, err = strconv.ParseUint(digits[2:], 8, 64)
			case strings.HasPrefix(digits, "0b"):
				unsigned, err = strconv.ParseUint(digits[2:], 2, 64)
			default:
				unsigned, err = strconv.ParseUint(digits, 10, 64)
			}

			if err != nil {
				return 0, errors.New(InvalidNumber, file, t)
			}

			return int(unsigned), nil
		}

		return int(signed), nil

	case token.Rune:
		r := t.Bytes(file.Bytes)
		r = unescape(r)

		if len(r) == 0 {
			return 0, errors.New(InvalidRune, file, t)
		}

		number, size := utf8.DecodeRune(r)

		if len(r) > size {
			return 0, errors.New(InvalidRune, file, t)
		}

		return int(number), nil
	}

	return 0, errors.New(InvalidNumber, file, t)
}