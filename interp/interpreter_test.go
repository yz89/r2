package interp

import (
	"r2/context"
	"testing"
)

func executeAdapter(exp string, env context.Env) (string, error) {
	return Execute(exp, env)
}

func Test_Execute(t *testing.T) {
	env0 := context.Env{}

	t.Log(R21)
	res1, err := executeAdapter(R21, env0)
	if err != nil {
		t.Error(err)
	}
	if res1 != "3" {
		t.Error("res wrong")
	}

	t.Log(R22)
	res2, err := executeAdapter(R22, env0)
	if err != nil {
		t.Error(err)
	}
	if res2 != "2" {
		t.Error("res wrong")
	}

	t.Log(R23)
	res3, err := executeAdapter(R23, env0)
	if err != nil {
		t.Error(err)
	}
	if res3 != "14" {
		t.Error("res wrong")
	}

	t.Log(R24)
	res4, err := executeAdapter(R24, env0)
	if err != nil {
		t.Error(err)
	}
	if res4 != "21" {
		t.Error("res wrong")
	}

	t.Log(R25)
	res5, err := executeAdapter(R25, env0)
	if err != nil {
		t.Error(err)
	}
	if res5 != "6" {
		t.Error("res wrong")
	}

	t.Log(R26)
	res6, err := executeAdapter(R26, env0)
	if err != nil {
		t.Error(err)
	}
	if res6 != "6" {
		t.Error("res wrong")
	}

	t.Log(R27)
	res7, err := executeAdapter(R27, env0)
	if err != nil {
		t.Error(err)
	}
	if res7 != "6" {
		t.Error("res wrong")
	}

	t.Log(R28)
	res8, err := executeAdapter(R28, env0)
	if err != nil {
		t.Error(err)
	}
	if res8 != "3" {
		t.Error("res wrong")
	}

	t.Log(R29)
	res9, err := executeAdapter(R29, env0)
	if err != nil {
		t.Error(err)
	}
	if res9 != "5" {
		t.Error("res wrong")
	}

	t.Log(R262)
	env1 := context.ExtEnv("x", "10", env0)
	env2 := context.ExtEnv("y", "3", env1)
	res62, err := executeAdapter(R262, env2)
	if err != nil {
		t.Error(err)
	}
	if res62 != "y;(* 4 y);{\"y\":\"3\"};{\"x\":\"10\"}" {
		t.Error("res wrong")
	}
}
