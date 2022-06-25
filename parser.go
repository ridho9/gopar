package gopar

type Parser func(input ParserInput) ParserResult

func (p Parser) Run(input string) ParserResult {
	return p(buildInput(input))
}

func Ref(p *Parser) Parser {
	return func(input ParserInput) ParserResult {
		return (*p)(input)
	}
}

type ParserResult struct {
	input       ParserInput
	result      any
	lexIdxStart int
	lexIdxEnd   int
	err         error
}

func (p ParserResult) Result() any {
	return p.result
}

func (p ParserResult) Input() ParserInput {
	return p.input
}

func (p ParserResult) Error() any {
	return p.err
}

func (p ParserResult) lex() string {
	return p.input.peekRange(p.lexIdxStart, p.lexIdxEnd)
}
