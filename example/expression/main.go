package main

import (
	"fmt"
)

// Implementation of the expression parser from
// http://www.craftinginterpreters.com/parsing-expressions.html.

func main() {
	input := `1 + 1 * -2 / 3`
	res := parser.Run(input)
	fmt.Printf("%v\nerr: %e\n", res.Result(), res.Error())
}
