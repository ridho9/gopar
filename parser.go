package gopar

// type Parser func(input string) (string, any, error)
type Parser struct {
	f func(input parserInput) ParserResult
}

func (p Parser) Run(input string) ParserResult {
	return p.f(buildInput(input))
}

type ParserResult struct {
	input  parserInput
	result any
	err    error
}

func (p ParserResult) Result() any {
	return p.result
}

func (p ParserResult) Error() any {
	return p.err
}
