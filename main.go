package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"r2/context"
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

	R21  = "(((+ 1 2)))"
	R22  = "(* 1 2)"
	R23  = "(* 2 (+ 3 4))"
	R24  = "(* (+ 1 2) (+ 3 4))"
	R25  = "((lambda (x) (* 2 x)) 3)"
	R26  = "(let ((x 2)) (let ((f (lambda (y) (* x y)))) (f 3)))"
	R261 = "(let ((f (lambda (y) (* 4 y)))) (f 3))"
	R262 = "(lambda (y) (* 4 y))"
	R27  = "(let ((x 2)) (let ((f (lambda (y) (* x y)))) (let ((x 4)) (f 3))))"
	R28  = "(let ((x 1)) (+ (let ((x 2)) x) x))"
	R29  = "(let ((x 1)) (let ((y 2)) (let ((x 3)) (+ x y))))"
)

type Closure struct {
	param string
	exp   string
	env   context.Env
}

func (c *Closure) Serialize() (string, error) {
	buf := c.param
	buf += ";"
	buf += c.exp
	for _, dict := range c.env {
		buf += ";"
		dictStr, err := json.Marshal(dict)
		if err != nil {
			return "", err
		}
		buf += string(dictStr)
	}

	return buf, nil
}

func (c *Closure) Deserialize(data string) error {
	dataArray := strings.Split(data, ";")
	c.param = dataArray[0]
	c.exp = dataArray[1]
	for _, dict := range dataArray[2:] {
		newMap := make(map[string]string)
		if err := json.Unmarshal([]byte(dict), &newMap); err != nil {
			return err
		}
		c.env = append(c.env, newMap)
	}
	return nil
}

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

func interp(exp string, env context.Env) (string, error) {
	fmt.Println("interp:", exp)

	// number
	if num, err := matchNumber(exp); err == nil {
		return num, nil
	}

	// variable
	if vari, err := matchSymbol(exp); err == nil {
		// find value in context
		v, err := context.Lookup(vari, env)
		if err != nil {
			return "", err
		}
		return v, nil
	}

	// match S expression
	if sExps, err := matchSexp(exp); err == nil {

		if len(sExps) == 2 {
			// call
			// fmt.Println("call", sExps)
			v1, err := interp(sExps[0], env)
			if err != nil {
				return "", err
			}
			v2, err := interp(sExps[1], env)
			if err != nil {
				return "", err
			}
			var cl Closure
			err = cl.Deserialize(v1)
			// fmt.Println("v1", v1)
			// fmt.Println("cl.env", cl.env)
			if err != nil {
				return "", err
			}
			newEnv := context.ExtEnv(cl.param, v2, cl.env)
			res, err := interp(cl.exp, newEnv)
			if err != nil {
				return "", err
			}
			return res, nil
		}

		op := sExps[0]
		switch op {
		case "lambda":
			// function
			// fmt.Println("function", sExps)
			cl := Closure{
				param: sExps[1][1 : len(sExps[1])-1],
				exp:   sExps[2],
				env:   env,
			}
			if clStr, err := cl.Serialize(); err == nil {
				return clStr, nil
			}
			return "", errors.New("error lambda exp")
		case "let":
			// bind
			// fmt.Println("bind", sExps)
			e, err := matchSexp(sExps[1])
			if err != nil {
				return "", err
			}
			x := e[0]
			v, err := interp(e[1], env)
			if err != nil {
				return "", err
			}
			newEnv := context.ExtEnv(x, v, env)
			res, err := interp(sExps[2], newEnv)
			if err != nil {
				return "", err
			}
			return res, nil
		default:
			// math operation
			v1, err := interp(sExps[1], env)
			if err != nil {
				return "", err
			}
			v2, err := interp(sExps[2], env)
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
	// exp, err := matchSexp(R262)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(exp)
	// for _, e := range exp {
	// 	fmt.Println(e)
	// }

	env0 := context.Env{}
	sum, err := interp(R27, env0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(sum)
}
