package interp

import (
	"errors"
	"strconv"
	"strings"
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

	// case (((((+ 1 2)))))
	if len(res) == 1 {
		res1, err := matchSexp(res[0])
		if err != nil {
			return nil, err
		}
		return res1, nil
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
