package gopar

import "errors"

type Parser func(input string) (string, any, error)

// Droplist(idx0, idx1,...) removes item with given indexes if parser returns a list,
// otherwise returns error.
func (p Parser) DropList(indexes ...int) Parser {
	return func(input string) (string, any, error) {
		input, res, err := p(input)
		if err != nil {
			return input, res, err
		}
		resL, ok := res.([]any)
		if !ok {
			return input, res, errors.New("result is not a list")
		}

		droppedResult := []any{}
		for idx, res := range resL {
			if len(indexes) == 0 {
				droppedResult = append(droppedResult, res)
				continue
			}
			if idx == indexes[0] {
				indexes = indexes[1:]
				continue
			}
			droppedResult = append(droppedResult, res)
			if idx > indexes[0] {
				indexes = indexes[1:]
			}
		}
		return input, droppedResult, nil
	}
}

func (p Parser) First() Parser {
	return func(input string) (string, any, error) {
		input, res, err := p(input)
		if err != nil {
			return input, res, err
		}
		resL, ok := res.([]any)
		if !ok {
			return input, res, errors.New("result is not a list")
		}
		if len(resL) == 0 {
			return input, res, errors.New("empty list")
		}

		return input, resL[0], nil
	}
}
