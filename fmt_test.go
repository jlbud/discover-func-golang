package main

import (
	"fmt"
	"testing"
)

func TestFmt(t *testing.T) {
	type t2 struct {
		C string
		D int
	}
	type t1 struct {
		A string
		B int
		t2
	}

	t11 := new(t1)
	t11.A = "a"
	t11.B = 1
	t11.C = "C"
	t11.D = 2

	t.Log(fmt.Sprintf("%+v", *t11))
}
