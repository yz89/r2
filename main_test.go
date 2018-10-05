package main

import (
	"testing"
)

func Test_R2(t *testing.T) {
	res1, err := calc(R21)
	if err != nil {
		t.Error(err)
	}
	if res1 != "3" {
		t.Error("res wrong")
	}

	res2, err := calc(R22)
	if err != nil {
		t.Error(err)
	}
	if res2 != "2" {
		t.Error("res wrong")
	}

	res3, err := calc(R23)
	if err != nil {
		t.Error(err)
	}
	if res3 != "14" {
		t.Error("res wrong")
	}

	res4, err := calc(R24)
	if err != nil {
		t.Error(err)
	}
	if res4 != "21" {
		t.Error("res wrong")
	}
}
