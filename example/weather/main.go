package main

import (
	"fmt"

	g "github.com/ridho9/gopar"
)

func main() {
	weatherString := g.String("Weather")
	timeString := g.Sequence(
		g.String("("),
		g.Or(
			g.String("today"),
			g.String("yesterday"),
			g.String("one week ago"),
		),
		g.String(")"),
	).DropList(0, 2).First()

	weatherType := g.Or(
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

	input, result, err := parser(input)

	fmt.Println(input)
	fmt.Printf("%#v\n", result)
	fmt.Println(err)
}
