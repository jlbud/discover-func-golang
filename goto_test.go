package main

import (
	"fmt"
	"testing"
)

func Test_goto(t *testing.T) {
	count := 0
retry:
	fmt.Println("a")
	count++

	if count == 100 {
		return
	}
	fmt.Println("c")

	for i := 0; i < 100; i++ {
		goto retry
	}
}

// 使用goto break最外层for循环
func TestGoto1(t *testing.T) {
	i := 0
f:
	for true {
		i++
		t.Logf("i: %d", i)
		if i == 10 {
			t.Logf("i is at %d break", i)
			break f
		}
	}
	t.Log("end")
}
