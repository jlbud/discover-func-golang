package main

import (
	"fmt"
	"testing"
)

var strs []string

func Test_main(t *testing.T) {
	if strs == nil {
		fmt.Println("strs is nil")
		return
	}

	strs = append(strs, "a")
	for _, s := range strs {
		fmt.Println(s)
	}
	fmt.Println("exit")
}
