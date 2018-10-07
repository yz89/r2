package interp

import (
	"errors"
	"strconv"
	"strings"
)

func MatchSexp(exp string) ([]string, error) {
	length := len(exp)
	if length < 2 {
		return nil, errors.New("length too short")
	}
	if exp[0] != '(' || exp[length-1] != ')' {
		return nil, errors.New("() not match")
	}

	type State byte
	const (
		S_VALUE State = iota
		S_SEXP
	)

	var res []string
	state := S_VALUE
	parCount := 0

	// remove the parenthesis
	exp = exp[1 : length-1]
	for i, c := range exp {
		switch state {
		case S_VALUE:
			if c == '(' {
				// the first (, enter S_SEXP
				parCount++
				res = append(res, "(")
				state = S_SEXP
			} else if c == ' ' {
				// skip space

			} else {
				// on S_VALUE, assume the first char of expression or the first char which followed space is the first char of a value
				if i == 0 || exp[i-1] == ' ' {
					// the first char of value
					res = append(res, string(c))
					continue
				}
				// read value, append c to the last string
				lastChar := []byte(res[len(res)-1])
				lastChar = append(lastChar, byte(c))
				res[len(res)-1] = string(lastChar)
			}
		case S_SEXP:
			if c == '(' {
				parCount++
			} else if c == ')' {
				parCount--
			}
			// read S expression, append c to the last string
			lastChar := []byte(res[len(res)-1])
			lastChar = append(lastChar, byte(c))
			res[len(res)-1] = string(lastChar)
			if parCount == 0 {
				// enter S_VALUE
				state = S_VALUE
			}
		default:
			panic("error state")
		}
	}

	// case (((((+ 1 2))))) -> (+ 1 2)
	if len(res) == 1 {
		resNext, err := MatchSexp(res[0])
		if err != nil {
			return nil, err
		}
		return resNext, nil
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
