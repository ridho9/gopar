package main

import (
	"fmt"

	g "github.com/ridho9/gopar"
)

func main() {
	weatherString := g.String("Weather")
	timeString := g.Sequence(
		g.String("("),
		g.Choice(
			g.String("today"),
			g.String("yesterday"),
			g.String("one week ago"),
		),
		g.String(")"),
	).TakeNth(1)

	weatherType := g.Choice(
		g.String("Sunny"),
		g.String("Cloudy"),
		g.String("Rainy"),
		g.String("Rain"),
	)

	parser := g.Sequence(
		weatherString,
		g.String(" "),
		timeString,
		g.String(": "),
		weatherType,
	)

	input := "Weather (today): Rainy"

	res := parser.Run(input)

	fmt.Printf("%#v\n", res.Result())
	fmt.Println(res.Error())
}
