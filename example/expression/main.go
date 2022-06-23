package main

import (
	"fmt"

	g "github.com/ridho9/gopar"
)

func main() {
	pNumber := g.TakeWhile(func(r rune) bool { return '0' <= r && r <= '9' })
	// Map(func(a any) any {
	// 	num, _ := strconv.ParseInt(a.(string), 10, 64)
	// 	return num
	// })

	pFloat := g.Sequence(
		pNumber,
		g.Optional(
			g.Sequence(
				g.String("."),
				pNumber,
			),
		),
	)

	parser := pFloat
	input := "123"
	res := parser.Run(input)
	fmt.Printf("%#v err: %e\n", res.Result(), res.Error())
}
