package gopar

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

// Matches exact string pattern. Parser result is string.
func String(pattern string) Parser {
	return func(input string) (string, any, error) {
		if len(input) == 0 {
			return input, "", ErrEndOfInput
		}

		if len(pattern) > len(input) {
			return input, "", errors.New("pattern longer than input")
		}

		iCursor := 0
		for pCursor, w := 0, 0; pCursor < len(pattern); pCursor += w {
			pRune, width := utf8.DecodeRuneInString(pattern[pCursor:])
			iRune, _ := utf8.DecodeRuneInString(input[iCursor:])
			if iRune == pRune {
				w = width
				iCursor += w
				continue
			} else {
				return input, "", fmt.Errorf(`expected "%s" found "%s"`, pattern, input)
			}
		}

		return input[iCursor:], input[:iCursor], nil
	}
}
