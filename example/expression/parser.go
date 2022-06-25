package main

import (
	"strconv"

	g "github.com/ridho9/gopar"
)

var (
	pNumber = g.TakeWhile(func(r rune) bool { return '0' <= r && r <= '9' })
	pFloat  = litWrap(g.Sequence(
		pNumber,
		g.Optional(
			g.Sequence(
				g.String("."),
				pNumber,
			),
		),
	).
		Recognize().
		Map(func(a any) any {
			num, _ := strconv.ParseFloat(a.(string), 64)
			return num
		}))

	pPrimary = g.Choice(pFloat)
	pUnary   g.Parser

	multispace0 = g.TakeWhile(func(r rune) bool { return r == ' ' })
	parser      g.Parser
)

func init() {
	pUnary = g.Choice(
		g.Sequence(
			g.Choice(opStr("+"), opStr("-")),
			g.Ref(&pUnary),
		).Map(func(a any) any {
			al := a.([]any)
			op := al[0]
			val := al[1].(float64)
			if op == "-" {
				val *= -1
			}
			return val
		}),
		pPrimary,
	)
	parser = pUnary
}

func litWrap(p g.Parser) g.Parser {
	return g.Delim(multispace0, p, multispace0)
}

func opStr(s string) g.Parser {
	return litWrap(g.String(s))
}
