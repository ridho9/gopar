package gopar

import (
	"errors"
	"fmt"
)

// Choice return the first result of matching parser
func Choice(parsers ...Parser) Parser {
	return Parser{
		f: func(input parserInput) (res ParserResult) {
			for _, parser := range parsers {
				res = parser.f(input)
				if res.err == nil {
					return res
				}
			}
			res.err = errors.New("no parser matches")
			return res
		},
	}
}

// Sequence return the results of running the parsers sequentially in a list.
// If any fail will return original input and error parser.
func Sequence(parsers ...Parser) Parser {
	return Parser{
		f: func(input parserInput) (res ParserResult) {
			origInput := input
			resultList := []any{}

			for _, parser := range parsers {
				res = parser.f(input)
				if res.err != nil {
					res.input = origInput
					res.err = fmt.Errorf("sequence fail: %w", res.err)
					return res
				}
				resultList = append(resultList, res.result)
				input = res.input
			}
			res.result = resultList
			return res
		},
	}
}

func Optional(parser Parser) Parser {
	return Parser{
		f: func(input parserInput) ParserResult {
			res := parser.f(input)
			if res.err != nil {
				res.result = nil
				res.err = nil
			}
			return res
		},
	}
}
