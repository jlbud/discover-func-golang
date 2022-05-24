package main

import (
	"fmt"
	"testing"
)

func TestGeneric(t *testing.T) {
	a := compare(30, 20)
	fmt.Println(a)
}

func compare[T int | int64](a, b T) bool {
	if a >= b {
		return true
	} else {
		return false
	}
}
