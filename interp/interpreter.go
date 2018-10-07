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
	t, exps, err := matchR2Type(exp)
	if err != nil {
		return "", err
	}

	switch t {
	case NUMBER:
		return exps[0], nil
	case BIND:
		bindPairs, err := MatchSexp(exps[0])
		if err != nil {
			return "", err
		}
		x := bindPairs[0]
		v, err := Execute(bindPairs[1], env)
		if err != nil {
			return "", err
		}
		newEnv := context.ExtEnv(x, v, env)
		res, err := Execute(exps[1], newEnv)
		if err != nil {
			return "", err
		}
		return res, nil
	case VAR:
		// find value in context
		v, err := context.Lookup(exps[0], env)
		if err != nil {
			return "", err
		}
		return v, nil
	case FUNCTION:
		var cl Closure
		cl.param = exps[0][1 : len(exps[0])-1]
		cl.exp = exps[1]
		cl.env = env
		res, err := cl.Serialize()
		if err != nil {
			return "", err
		}
		return res, nil
	case CALL:
		f, err := Execute(exps[0], env)
		if err != nil {
			return "", err
		}
		arg, err := Execute(exps[1], env)
		if err != nil {
			return "", err
		}
		var cl Closure
		if err := cl.Deserialize(f); err != nil {
			return "", err
		}

		newEnv := context.ExtEnv(cl.param, arg, cl.env)
		res, err := Execute(cl.exp, newEnv)
		if err != nil {
			return "", err
		}
		return res, nil
	case OPERATION:
		op := exps[0]
		v1, err := Execute(exps[1], env)
		if err != nil {
			return "", err
		}
		v2, err := Execute(exps[2], env)
		if err != nil {
			return "", err
		}
		res, err := strOperation(op, v1, v2)
		if err != nil {
			return "", err
		}
		return res, nil
	default:
	}

	return "", errors.New("executed failed")
}
