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
	).Recognize()

	parser := pFloat
	input := "123.45"
	res := parser.Run(input)
	fmt.Printf("%#v\nerr: %e\n", res.Result(), res.Error())
}
