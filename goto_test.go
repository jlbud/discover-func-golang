package main

import (
	"fmt"
	"testing"
)

func Test_goto(t *testing.T) {
	count := 0
retry:
	fmt.Println("a")
	count ++

	if count == 100 {
		return
	}
	fmt.Println("c")

	for i := 0; i < 100; i ++ {
		goto retry
	}
}
