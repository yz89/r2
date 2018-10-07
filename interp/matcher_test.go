package interp

import (
	"testing"
)

func matchAdapter(exp string) ([]string, error) {
	return MatchSexp(exp)
}

func Test_MatchSexp(t *testing.T) {

	t.Log(R21)
	exp1, err := matchAdapter(R21)
	if err != nil {
		t.Error(err)
	}
	if exp1[0] != "+" || exp1[1] != "1" || exp1[2] != "2" {
		t.Error("error match", exp1)
	}

	t.Log(R22)
	exp2, err := matchAdapter(R22)
	if err != nil {
		t.Error(err)
	}
	if exp2[0] != "*" || exp2[1] != "1" || exp2[2] != "2" {
		t.Error("error match")
	}

	t.Log(R23)
	exp3, err := matchAdapter(R23)
	if err != nil {
		t.Error(err)
	}
	if exp3[0] != "*" || exp3[1] != "2" || exp3[2] != "(+ 3 4)" {
		t.Error("error match")
	}

	t.Log(R24)
	exp4, err := matchAdapter(R24)
	if err != nil {
		t.Error(err)
	}
	if exp4[0] != "*" || exp4[1] != "(+ 1 2)" || exp4[2] != "(+ 3 4)" {
		t.Error("error match")
	}

	t.Log(R25)
	exp5, err := matchAdapter(R25)
	if err != nil {
		t.Error(err)
	}
	if exp5[0] != "(lambda (x) (* 2 x))" || exp5[1] != "3" {
		t.Error("error match")
	}

	t.Log(R26)
	exp6, err := matchAdapter(R26)
	if err != nil {
		t.Error(err)
	}
	if exp6[0] != "let" || exp6[1] != "((x 2))" || exp6[2] != "(let ((f (lambda (y) (* x y)))) (f 3))" {
		t.Error("error match")
	}

	t.Log(R261)
	exp61, err := matchAdapter(R261)
	if err != nil {
		t.Error(err)
	}
	if exp61[0] != "let" || exp61[1] != "((f (lambda (y) (* 4 y))))" || exp61[2] != "(f 3)" {
		t.Error("error match")
	}

	t.Log(R262)
	exp62, err := matchAdapter(R262)
	if err != nil {
		t.Error(err)
	}
	if exp62[0] != "lambda" || exp62[1] != "(y)" || exp62[2] != "(* 4 y)" {
		t.Error("error match")
	}

	t.Log(R27)
	exp27, err := matchAdapter(R27)
	if err != nil {
		t.Error(err)
	}
	if exp27[0] != "let" || exp27[1] != "((x 2))" || exp27[2] != "(let ((f (lambda (y) (* x y)))) (let ((x 4)) (f 3)))" {
		t.Error("error match")
	}

	t.Log(R28)
	exp28, err := matchAdapter(R28)
	if err != nil {
		t.Error(err)
	}
	if exp28[0] != "let" || exp28[1] != "((x 1))" || exp28[2] != "(+ (let ((x 2)) x) x)" {
		t.Error("error match")
	}

	t.Log(R29)
	exp29, err := matchAdapter(R29)
	if err != nil {
		t.Error(err)
	}
	if exp29[0] != "let" || exp29[1] != "((x 1))" || exp29[2] != "(let ((y 2)) (let ((x 3)) (+ x y)))" {
		t.Error("error match")
	}
}
