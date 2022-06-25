package gopar

import "unicode/utf8"

type ParserInput struct {
	text      string
	textLen   int
	cursor    int
	spanStart int
}

func buildInput(text string) ParserInput {
	return ParserInput{
		text:      text,
		textLen:   len(text),
		cursor:    0,
		spanStart: 0,
	}
}

func (pi ParserInput) len() int {
	return pi.textLen - pi.cursor
}

func (pi ParserInput) peekRune() rune {
	r, _ := utf8.DecodeRuneInString(pi.text[pi.cursor:])
	return r
}

func (pi *ParserInput) popRune() (rune, int) {
	r, w := utf8.DecodeRuneInString(pi.text[pi.cursor:])
	pi.cursor += w
	return r, w
}

func (pi *ParserInput) rwdCursor(dif int) {
	pi.cursor -= dif
}

func (pi *ParserInput) takeSpan() string {
	res := pi.text[pi.spanStart:pi.cursor]
	pi.spanStart = pi.cursor
	return res
}

func (pi *ParserInput) peekStringLen(len int) string {
	return pi.text[pi.cursor : pi.cursor+len]
}

func (pi *ParserInput) peekString() string {
	return pi.text[pi.cursor:]
}

func (pi ParserInput) peekRange(startIdx int, endIdx int) string {
	return pi.text[startIdx:endIdx]
}
