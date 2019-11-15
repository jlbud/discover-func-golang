package main

import (
	"fmt"
	"testing"
)

func TestSwitch(t *testing.T) {
	v := 1
	switch v {
	case 1, 5:
		fmt.Println("case is: ", v)
		return
	}
	fmt.Println("success")
}

func TestSwitch1(t *testing.T) {
	switch 1 {
	case 1, 2:
	default:
		fmt.Println("abcd")
	}
}
