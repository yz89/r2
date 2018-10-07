package interp

import (
	"errors"
	"fmt"
	"r2/context"
	"strconv"
)

const (
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

func Execute(exp string, env context.Env) (string, error) {
	fmt.Println("Execute:", exp)

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
	if sExps, err := MatchSexp(exp); err == nil {

		if len(sExps) == 2 {
			// function call
			// fmt.Println("call", sExps)
			v1, err := Execute(sExps[0], env)
			if err != nil {
				return "", err
			}
			v2, err := Execute(sExps[1], env)
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
			res, err := Execute(cl.exp, newEnv)
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
			e, err := MatchSexp(sExps[1])
			if err != nil {
				return "", err
			}
			x := e[0]
			v, err := Execute(e[1], env)
			if err != nil {
				return "", err
			}
			newEnv := context.ExtEnv(x, v, env)
			res, err := Execute(sExps[2], newEnv)
			if err != nil {
				return "", err
			}
			return res, nil
		default:
			// math operation
			v1, err := Execute(sExps[1], env)
			if err != nil {
				return "", err
			}
			v2, err := Execute(sExps[2], env)
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
