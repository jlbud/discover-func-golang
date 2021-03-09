package main

import (
	"fmt"
	"testing"
)

func TestDefer1(t *testing.T) {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

/**
10 1 2 3 // 先执行 calc("10", a, b)
20 0 2 2 // 执行  calc("20", a, b)
2 0 2 2 执行 calc("2", a, calc("20", a, b))
1 1 3 4 执行 calc("1", a, calc("10", a, b))，此时还是a=1
*/
func TestDefer2(t *testing.T) {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
}
