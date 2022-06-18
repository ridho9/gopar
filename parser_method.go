package gopar

import (
	"errors"
	"sort"
)

// Droplist(idx0, idx1,...) removes item with given indexes if parser returns a list,
// otherwise returns error.
func (p Parser) DropList(indexes ...int) Parser {
	sort.Slice(indexes, func(i, j int) bool {
		return i < j
	})
	return Parser{
		f: func(input parserInput) ParserResult {
			res := p.f(input)
			if res.err != nil {
				return res
			}
			resL, ok := res.result.([]any)
			if !ok {
				res.err = errors.New("result is not a list")
				return res
			}

			droppedResult := []any{}
			for idx, res := range resL {
				if len(indexes) == 0 {
					droppedResult = append(droppedResult, res)
					continue
				}
				if idx == indexes[0] {
					indexes = indexes[1:]
					continue
				}
				droppedResult = append(droppedResult, res)
				if idx > indexes[0] {
					indexes = indexes[1:]
				}
			}
			res.result = droppedResult
			return res
		},
	}
}

// Returns the first item if the result is a list, else error.
func (p Parser) First() Parser {
	return Parser{
		f: func(input parserInput) ParserResult {
			res := p.f(input)
			if res.err != nil {
				return res
			}
			resL, ok := res.result.([]any)
			if !ok {
				res.err = errors.New("result is not a list")
				return res
			}
			if len(resL) == 0 {
				res.err = errors.New("empty list")
				return res
			}

			res.result = resL[0]
			return res
		},
	}
}
