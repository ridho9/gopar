package gopar

import "unicode/utf8"

type parserInput struct {
	text      string
	textLen   int
	cursor    int
	spanStart int
}

func buildInput(text string) parserInput {
	return parserInput{
		text:      text,
		textLen:   len(text),
		cursor:    0,
		spanStart: 0,
	}
}

func (pi parserInput) len() int {
	return pi.textLen - pi.cursor
}

func (pi parserInput) peekRune() rune {
	r, _ := utf8.DecodeRuneInString(pi.text[pi.cursor:])
	return r
}

func (pi *parserInput) popRune() (rune, int) {
	r, w := utf8.DecodeRuneInString(pi.text[pi.cursor:])
	pi.cursor += w
	return r, w
}

func (pi *parserInput) rwdCursor(dif int) {
	pi.cursor -= dif
}

func (pi *parserInput) takeSpan() string {
	res := pi.text[pi.spanStart:pi.cursor]
	pi.spanStart = pi.cursor
	return res
}

func (pi *parserInput) peekStringLen(len int) string {
	return pi.text[pi.cursor : pi.cursor+len]
}

func (pi *parserInput) peekString() string {
	return pi.text[pi.cursor:]
}

func (pi parserInput) peekRange(startIdx int, endIdx int) string {
	return pi.text[startIdx:endIdx]
}
