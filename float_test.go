package main

import "testing"

func Test_float(t *testing.T) {
	var a float64
	var b float64

	a = 0
	b = 0.0
	if a == b {
		t.Log("equal")
		return
	}
	t.Log("return")
}
