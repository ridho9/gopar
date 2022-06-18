package gopar

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

// Matches exact string pattern. Parser result is string.
func String(pattern string) Parser {
	return Parser{
		f: func(input parserInput) (res ParserResult) {
			res.input = input
			if input.len() == 0 {
				res.err = ErrEndOfInput
				return res
			}

			if len(pattern) > input.len() {
				res.err = errors.New("pattern longer than input")
				return res
			}

			for pCursor, w := 0, 0; pCursor < len(pattern); pCursor += w {
				pRune, width := utf8.DecodeRuneInString(pattern[pCursor:])
				iRune, _ := input.peekRune()
				if iRune == pRune {
					w = width
					input.advCursor(w)
					continue
				} else {
					res.err = fmt.Errorf(`expected "%s" found "%s"`, pattern, input.peekStringLen(len(pattern)))
					return res
				}
			}

			res.result = input.takeSpan()
			res.input = input
			return res
		},
	}
}
