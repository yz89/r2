package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	SUM1 = "(1 2)"
	SUM2 = "(1 (2 3))"
	SUM3 = "((1 2) 3)"
	SUM4 = "((1 2) (3 4))"
	SUM5 = "((1 2) (3 4) (5 6) (7 8))"

	CALC1 = "(+ 1 2)"
	CALC2 = "(* 1 2)"
	CALC3 = "(* (+ 1 2) (+ 3 4))"

	R21 = "(+ 1 2)"
	R22 = "(* 1 2)"
	R23 = "(* 2 (+ 3 4))"
	R24 = "(* (+ 1 2) (+ 3 4))"
	R25 = "((lambda (x) (* 2 x)) 3)"
	R26 = "(let ((x 2)) (let ((f (lambda (y) (* x y)))) (f 3)))"
	R27 = "(let ((x 2)) (let ((f (lambda (y) (* x y)))) (let ((x 4)) (f 3))))"
)

func matchSexp(exp string) ([]string, error) {
	length := len(exp)
	if length < 2 {
		return nil, errors.New("length too short")
	}
	if exp[0] != '(' || exp[length-1] != ')' {
		return nil, errors.New("() not match")
	}

	var res []string
	q := 0
	exp = exp[1 : length-1]
	for i, c := range exp {
		if c == ')' {
			q--
		}

		if c == '(' {
			if q == 0 {
				// the first (
				q++
				res = append(res, "(")
				continue
			}
			q++
			lastElement := []byte(res[len(res)-1])
			lastElement = append(lastElement, byte(c))
			res[len(res)-1] = string(lastElement)
			continue
		} else if q > 0 {
			lastElement := []byte(res[len(res)-1])
			lastElement = append(lastElement, byte(c))
			res[len(res)-1] = string(lastElement)
			continue
		}

		if c == ' ' && q == 0 {
			continue
		}

		if q == 0 {
			if i == 0 || exp[i-1] == ' ' {
				// the first char
				res = append(res, string(c))
				continue
			}
			lastElement := []byte(res[len(res)-1])
			lastElement = append(lastElement, byte(c))
			res[len(res)-1] = string(lastElement)
		}
	}

	return res, nil
}

func matchNumber(exp string) (string, error) {
	_, err := strconv.Atoi(exp)
	if err != nil {
		return "", err
	}

	return exp, nil
}

func matchSymbol(exp string) (string, error) {
	ok := strings.ContainsAny(exp, "() ")
	if ok {
		return "", errors.New("not a symbol")
	}
	return exp, nil
}

func strOperation(op, v1, v2 string) (string, error) {
	num1, err := strconv.Atoi(v1)
	if err != nil {
		return "", err
	}
	num2, err := strconv.Atoi(v2)
	if err != nil {
		return "", err
	}

	var res int
	switch op {
	case "+":
		res = num1 + num2
	case "-":
		res = num1 - num2
	case "*":
		res = num1 * num2
	case "/":
		res = num1 / num2
	default:
		return "", errors.New("invalid operation")
	}

	resStr := strconv.Itoa(res)
	return resStr, nil
}

func calc(exp string) (string, error) {
	fmt.Println("calc : ", exp)

	// number
	if num, err := matchNumber(exp); err == nil {
		return num, nil
	}
	// symbol
	if symbol, err := matchSymbol(exp); err == nil {
		// find value in context
		return symbol, nil
	}

	// match S expression
	if sExps, err := matchSexp(exp); err == nil {
		op := sExps[0]

		switch op {
		case "lambda":
			// function

		case "let":
			// bind

		default:
			// math operation
			v1, err := calc(sExps[1])
			if err != nil {
				return "", err
			}
			v2, err := calc(sExps[2])
			if err != nil {
				return "", err
			}

			sum, err := strOperation(op, v1, v2)
			if err != nil {
				return "", err
			}
			return sum, nil
		}
	}

	return "", errors.New("error match")
}

func main() {
	exp, err := matchSexp(R25)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(exp)
	for _, e := range exp {
		fmt.Println(e)
	}

	// num, err := matchNumber("123dd")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(num)

	// sum, err := calc(R25)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(sum)
}
