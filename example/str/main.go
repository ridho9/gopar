package main

import (
	"fmt"
	"strings"

	g "github.com/ridho9/gopar"
)

func main() {
	parser := g.Delim(
		g.String(`"`),
		g.Many0(
			g.Choice(
				g.String(`\n`).Value("\n"),
				g.String(`\"`).Value(`"`),
				g.String(`\\`).Value(`\`),
				g.TakeWhile1(func(r rune) bool {
					return r != '"'
				}),
			),
		).Map(func(a any) any {
			al := a.([]any)
			b := strings.Builder{}
			for _, l := range al {
				b.WriteString(l.(string))
			}
			return b.String()
		}),
		g.String(`"`),
	)

	input := `"\n\"\\n \tasdasda日本日本語sdasdasda"`

	res := parser.Run(input)
	fmt.Printf("[%#v]\nerr: %e\n", res.Result(), res.Error())
}
