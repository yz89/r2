package main

import (
	"errors"
	"fmt"
	"strconv"
)

const (
	SUM1 = "(1 2)"
	SUM2 = "(1 (2 3))"
	SUM3 = "((1 2) 3)"
	SUM4 = "((1 2) (3 4))"
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

		if c == '(' && q == 0 {
			// the first (
			q++
			subExp := "("
			res = append(res, subExp)
			continue
		} else if q > 0 {
			lastElement := []byte(res[len(res)-1])
			lastElement = append(lastElement, byte(c))
			res[len(res)-1] = string(lastElement)
			continue
		}

		if c == ' ' {
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

	// res must only have two elements
	return res, nil
}

func matchNumber(exp string) (string, error) {
	_, err := strconv.Atoi(exp)
	if err != nil {
		return "", err
	}

	return exp, nil
}

func strSum(v1 string, v2 string) (string, error) {
	num1, err := strconv.Atoi(v1)
	if err != nil {
		return "", err
	}

	num2, err := strconv.Atoi(v2)
	if err != nil {
		return "", err
	}

	sum := num1 + num2
	sumStr := strconv.Itoa(sum)
	return sumStr, nil
}

func treeSum(exp string) (string, error) {
	fmt.Println("tree sum : ", exp)

	// match number
	if num, err := matchNumber(exp); err == nil {
		return num, nil
	}

	// match S expression
	if sExps, err := matchSexp(exp); err == nil {
		v1, err := treeSum(sExps[0])
		if err != nil {
			return "", err
		}
		v2, err := treeSum(sExps[1])
		if err != nil {
			return "", err
		}
		sum, err := strSum(v1, v2)
		if err != nil {
			return "", err
		}
		return sum, nil
	}

	return "", errors.New("error match")
}

func main() {
	// exp, err := matchSexp(SUM4)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(exp)

	// num, err := matchNumber("123dd")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(num)

	sum, err := treeSum(SUM4)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(sum)
}
