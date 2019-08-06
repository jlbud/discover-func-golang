package main

import (
	"fmt"
	"testing"
)

func TestInterfaceNil(t *testing.T) {
	a := &A{
	}
	fmt.Println(a)
	isNil(nil)
}

func isNil(param interface{}) {
	a, ok := param.(*A)
	if !ok || a == nil {
		fmt.Println("ok value is ", ok)
		fmt.Println("this is not ok ", a)
		return
	}

	fmt.Println("this is ok ", a)
}
