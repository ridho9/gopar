package main

import (
	"fmt"
)

func main() {
	input := ` + 123.45 `
	res := parser.Run(input)
	fmt.Printf("%#v\nerr: %e\n", res.Result(), res.Error())
}
