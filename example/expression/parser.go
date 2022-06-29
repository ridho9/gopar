package main

import (
	"strconv"

	g "github.com/ridho9/gopar"
)

var (
	pNumber  g.Parser
	pFloat   g.Parser
	pPrimary g.Parser
	pUnary   g.Parser
	pFactor  g.Parser
	pTerm    g.Parser

	multispace0 = g.TakeWhile0(func(r rune) bool { return r == ' ' })
	parser      g.Parser
)

func init() {
	pNumber = g.TakeWhile1(func(r rune) bool { return '0' <= r && r <= '9' })
	pFloat = litWrap(g.Sequence(
		pNumber,
		g.Optional(
			g.Sequence(
				g.String("."),
				pNumber,
			),
		)).
		Recognize().
		Map(func(a any) any {
			num, _ := strconv.ParseFloat(a.(string), 64)
			return Literal{Value: num}
		}))

	pPrimary = g.Choice(pFloat)
	pUnary = g.Choice(
		g.Sequence(
			g.Choice(opStr("!"), opStr("-")),
			g.Ref(&pUnary),
		).Map(func(a any) any {
			al := a.([]any)
			return Unary{Op: al[0].(string), Value: al[1]}
		}),
		pPrimary,
	)

	pFactor = g.Sequence(
		pUnary,
		g.Many0(
			g.Sequence(
				g.Choice(opStr("*"), opStr("/")),
				pUnary,
			),
		),
	).Map(packBinOp)

	pTerm = g.Sequence(
		pFactor,
		g.Many0(
			g.Sequence(
				g.Choice(opStr("+"), opStr("-")),
				pFactor,
			),
		),
	).Map(packBinOp)

	parser = g.Sequence(pTerm, g.EOF()).First()
}

func litWrap(p g.Parser) g.Parser {
	return g.Delim(multispace0, p, multispace0)
}

func opStr(s string) g.Parser {
	return litWrap(g.String(s))
}

func packBinOp(a any) any {
	al := a.([]any)
	left := al[0]
	reps := al[1].([]any)
	if len(reps) == 0 {
		return left
	} else {
		for _, repItem := range reps {
			repL := repItem.([]any)
			op, right := repL[0].(string), repL[1]
			newTree := BinaryOp{LeftVal: left, Op: op, RightVal: right}
			left = newTree
		}
		return left
	}
}
