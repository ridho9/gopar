package main

import "fmt"

type Literal struct {
	Value any
}

func (l Literal) String() string {
	return fmt.Sprintf("%v", l.Value)
}

type Unary struct {
	Op    string
	Value any
}

func (u Unary) String() string {
	return fmt.Sprintf("%s%v", u.Op, u.Value)
}

type BinaryOp struct {
	LeftVal  any
	Op       string
	RightVal any
}

func (b BinaryOp) String() string {
	return fmt.Sprintf("(%v %v %v)", b.LeftVal, b.Op, b.RightVal)
}
