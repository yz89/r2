package main

import (
	"r2/context"
	"testing"
)

func Test_R2(t *testing.T) {
	env0 := context.Env{}

	res1, err := interp(R21, env0)
	if err != nil {
		t.Error(err)
	}
	if res1 != "3" {
		t.Error("res wrong")
	}

	res2, err := interp(R22, env0)
	if err != nil {
		t.Error(err)
	}
	if res2 != "2" {
		t.Error("res wrong")
	}

	res3, err := interp(R23, env0)
	if err != nil {
		t.Error(err)
	}
	if res3 != "14" {
		t.Error("res wrong")
	}

	res5, err := interp(R25, env0)
	if err != nil {
		t.Error(err)
	}
	if res5 != "6" {
		t.Error("res wrong")
	}

	res6, err := interp(R26, env0)
	if err != nil {
		t.Error(err)
	}
	if res6 != "6" {
		t.Error("res wrong")
	}

	res7, err := interp(R27, env0)
	if err != nil {
		t.Error(err)
	}
	if res7 != "6" {
		t.Error("res wrong")
	}

	res8, err := interp(R28, env0)
	if err != nil {
		t.Error(err)
	}
	if res8 != "3" {
		t.Error("res wrong")
	}

	res9, err := interp(R29, env0)
	if err != nil {
		t.Error(err)
	}
	if res9 != "5" {
		t.Error("res wrong")
	}

	env1 := context.ExtEnv("x", "10", env0)
	env2 := context.ExtEnv("y", "3", env1)
	res62, err := interp(R262, env2)
	if err != nil {
		t.Error(err)
	}
	if res62 != "y;(* 4 y);{\"y\":\"3\"};{\"x\":\"10\"}" {
		t.Error("res wrong")
	}
}
