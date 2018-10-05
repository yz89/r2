package context

import (
	"testing"
)

func TestContext_ExtEnv(t *testing.T) {
	x1 := "name"
	v1 := "andrew"
	env0 := Env{}
	env1 := ExtEnv(x1, v1, env0)

	x2 := "age"
	v2 := "30"
	env2 := ExtEnv(x2, v2, env1)

	x3 := "gender"
	v3 := "male"
	env3 := ExtEnv(x3, v3, env2)

	if len(env1) != 1 || env1[0]["name"] != "andrew" {
		t.Error("error env1", env1)
	}

	if len(env2) != 2 || env2[0]["age"] != "30" || env2[1]["name"] != "andrew" {
		t.Error("error env2", env2)
	}

	if len(env3) != 3 || env3[0]["gender"] != "male" || env3[1]["age"] != "30" || env3[2]["name"] != "andrew" {
		t.Error("error env3", env3)
	}

	t.Log(env1, env2, env3)
}

func TestContext_Lookup(t *testing.T) {
	x1 := "x"
	v1 := "10"
	env0 := Env{}
	env1 := ExtEnv(x1, v1, env0)

	x2 := "y"
	v2 := "4"
	env2 := ExtEnv(x2, v2, env1)

	x3 := "x"
	v3 := "6"
	env3 := ExtEnv(x3, v3, env2)

	// env0
	if _, err := Lookup("x", env0); err == nil {
		t.Error("x shoud not be found in env0")
	}

	// env1
	xv1, err := Lookup("x", env1)
	if err != nil {
		t.Error("x shoud be found in env1", err)
	}
	if xv1 != "10" {
		t.Error("x shoud be equal to 10", xv1)
	}

	if _, err := Lookup("y", env1); err == nil {
		t.Error("y shoud not be found in env1")
	}

	// env2
	xv2, err := Lookup("x", env2)
	if err != nil {
		t.Error("x shoud be found in env2", err)
	}
	if xv2 != "10" {
		t.Error("x shoud be equal to 10", xv2)
	}

	yv2, err := Lookup("y", env2)
	if err != nil {
		t.Error("y shoud be found in env2", err)
	}
	if yv2 != "4" {
		t.Error("y shoud be equal to 4", yv2)
	}

	// env3
	xv3, err := Lookup("x", env3)
	if err != nil {
		t.Error("x shoud be found in env3", err)
	}
	if xv3 != "6" {
		t.Error("x shoud be equal to 6", xv3)
	}

	yv3, err := Lookup("y", env3)
	if err != nil {
		t.Error("y shoud be found in env3", err)
	}
	if yv3 != "4" {
		t.Error("y shoud be equal to 4", yv3)
	}
}
