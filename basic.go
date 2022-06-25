package gopar

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

// Matches exact string pattern. Parser result is string.
func String(pattern string) Parser {
	return func(input ParserInput) (res ParserResult) {
		res.lexIdxStart = input.cursor
		res.lexIdxEnd = input.cursor
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
			iRune, iW := input.popRune()
			if iRune == pRune {
				w = width
				continue
			} else {
				input.rwdCursor(iW)
				res.err = fmt.Errorf(`expected "%s" found "%s..."`, pattern, input.peekStringLen(len(pattern)))
				return res
			}
		}

		res.result = input.takeSpan()
		res.input = input
		res.lexIdxEnd = input.cursor
		return res
	}
}

func TakeWhile0(pred func(rune) bool) Parser {
	return func(input ParserInput) (res ParserResult) {
		res.lexIdxStart = input.cursor
		res.lexIdxEnd = input.cursor
		for {
			if input.len() == 0 {
				break
			}

			iRune, iW := input.popRune()
			if !pred(iRune) {
				input.rwdCursor(iW)
				break
			}
		}
		res.result = input.takeSpan()
		res.input = input
		res.lexIdxEnd = input.cursor
		return res
	}
}

func TakeWhile1(pred func(rune) bool) Parser {
	return func(input ParserInput) ParserResult {
		res := TakeWhile0(pred)(input)
		if res.err != nil {
			return res
		}
		if len(res.result.(string)) == 0 {
			res.err = fmt.Errorf("error TakeWhile1")
		}
		return res
	}
}
