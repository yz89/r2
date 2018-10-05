package main

import (
	"errors"
	"fmt"
)

const (
	SUM1 = "(1 2)"
	SUM2 = "(1 (2 3))"
	SUM3 = "((1 2) 3)"
	SUM4 = "((1 2) (3 4))"
)

func matchSexp(exp string) (string, error) {
	length := len(exp)
	if length < 2 {
		return "", errors.New("length too short")
	}
	if exp[0] != '(' || exp[length-1] != ')' {
		return "", errors.New("() not match")
	}

	return exp[1 : length-1], nil
}

func main() {
	exp, err := matchSexp(SUM4)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(exp)
}
