package gopar

import (
	"errors"
	"fmt"
)

// Or return the first result of matching parser
func Or(parsers ...Parser) Parser {
	return func(input string) (nextInput string, result any, err error) {
		for _, parser := range parsers {
			nextInput, result, err = parser(input)
			if err == nil {
				return nextInput, result, err
			}
		}
		return input, result, errors.New("no parser matches")
	}
}

// Sequence return the results of running the parsers sequentially in a list.
// If any fail will return original input and error parser.
func Sequence(parsers ...Parser) Parser {
	return func(input string) (string, any, error) {
		var err error
		var pRes any
		origInput := input
		resultList := []any{}

		for _, parser := range parsers {
			input, pRes, err = parser(input)
			if err != nil {
				return origInput, []any{}, fmt.Errorf("sequence fail: %w", err)
			}
			resultList = append(resultList, pRes)
		}
		return input, resultList, err
	}
}
