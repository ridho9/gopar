package gopar

// type Parser func(input string) (string, any, error)
type Parser struct {
	fn func(input parserInput) ParserResult
}

func (p Parser) Run(input string) ParserResult {
	return p.fn(buildInput(input))
}

func Ref(p *Parser) Parser {
	return Parser{
		fn: func(input parserInput) ParserResult {
			return p.fn(input)
		},
	}
}

type ParserResult struct {
	input       parserInput
	result      any
	lexIdxStart int
	lexIdxEnd   int
	err         error
}

func (p ParserResult) Result() any {
	return p.result
}

func (p ParserResult) Error() any {
	return p.err
}

func (p ParserResult) Lex() string {
	return p.input.peekRange(p.lexIdxStart, p.lexIdxEnd)
}
