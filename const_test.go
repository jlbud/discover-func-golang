package main

import (
	"fmt"
	"testing"
)

func TestConst(t *testing.T) {
	const (
		a = 7
		b
		c
	)
	fmt.Println(a, b, c)
}
