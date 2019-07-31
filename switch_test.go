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
	}
}
