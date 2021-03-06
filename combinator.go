package gopar

import (
	"errors"
	"fmt"
)

// Choice return the first result of matching parser
func Choice(parsers ...Parser) Parser {
	return func(input ParserInput) (res ParserResult) {
		for _, parser := range parsers {
			res = parser(input)
			if res.err == nil {
				return res
			}
		}
		res.err = errors.New("no parser matches")
		return res
	}
}

// Sequence return the results of running the parsers sequentially in a list.
// If any fail will return original input and error parser.
func Sequence(parsers ...Parser) Parser {
	return func(input ParserInput) (res ParserResult) {
		origInput := input
		resultList := []any{}
		lexIdxStart := origInput.cursor

		for _, parser := range parsers {
			res = parser(input)
			if res.err != nil {
				res.input = origInput
				res.err = fmt.Errorf("sequence fail: %w", res.err)
				return res
			}
			resultList = append(resultList, res.result)
			input = res.input
		}
		res.result = resultList
		res.lexIdxStart = lexIdxStart
		return res
	}
}

func Optional(parser Parser) Parser {
	return func(input ParserInput) ParserResult {
		res := parser(input)
		if res.err != nil {
			res.result = nil
			res.err = nil
		}
		return res
	}
}

func Delim(p1 Parser, p2 Parser, p3 Parser) Parser {
	return Sequence(p1, p2, p3).TakeNth(1)
}

func Many0(p Parser) Parser {
	return func(input ParserInput) (res ParserResult) {
		lexIdxStart := input.cursor
		resultList := []any{}
		for {
			res = p(input)
			if res.err != nil {
				res.err = nil
				break
			}
			resultList = append(resultList, res.result)
			input = res.input
		}
		res.result = resultList
		res.lexIdxStart = lexIdxStart
		return res
	}
}

func EOF() Parser {
	return func(input ParserInput) (res ParserResult) {
		res.input = input
		if input.cursor < input.textLen {
			res.err = fmt.Errorf("expected end of input")
			return res
		}
		return res
	}
}
