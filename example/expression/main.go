package main

import (
	"fmt"
	"strconv"

	g "github.com/ridho9/gopar"
)

var (
	pNumber = g.TakeWhile(func(r rune) bool { return '0' <= r && r <= '9' })

	pFloat = g.Sequence(
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
		})

	multispace0 = g.TakeWhile(func(r rune) bool { return r == ' ' })
)

func main() {
	parser := ignoreWs(pFloat)
	input := ` 123.45 `
	res := parser.Run(input)
	fmt.Printf("%#v\nerr: %e\n", res.Result(), res.Error())
}

func ignoreWs(p g.Parser) g.Parser {
	return g.Delim(multispace0, p, multispace0)
}
