package gopar

import (
	"errors"
)

// Returns the first item if the result is a list, else error.
func (p Parser) First() Parser {
	return func(input ParserInput) ParserResult {
		res := p(input)
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
	}
}

func (p Parser) Map(mapper func(any) any) Parser {
	return func(input ParserInput) ParserResult {
		res := p(input)
		if res.err != nil {
			return res
		}
		res.result = mapper(res.result)
		return res
	}
}

func (p Parser) TakeNth(n int) Parser {
	return p.Map(func(a any) any {
		al := a.([]any)
		return al[n]
	})
}

func (p Parser) Recognize() Parser {
	return func(input ParserInput) ParserResult {
		res := p(input)
		if res.err != nil {
			return res
		}
		res.result = res.lex()
		return res
	}
}
