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
	).DropList(0, 2).First()

	weatherType := g.Choice(
		g.String("Sunny"),
		g.String("Cloudy"),
		g.String("Rain"),
	)

	parser := g.Sequence(
		weatherString,
		g.String(" "),
		timeString,
		g.String(": "),
		weatherType,
	).DropList(1, 3)

	input := "Weather (today): Sunny"

	res := parser.Run(input)

	fmt.Printf("%#v\n", res.Result())
	fmt.Println(res.Error())
}
