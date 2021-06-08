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

var _ O = new(Boy)

type O interface {
	walk()
	run()
}
type S struct {
}

func (_ Man) walk() {
	fmt.Println("walk")
}

func (_ Man) run() {
	fmt.Println("run")
}

type Boy struct {
	Man
}

func TestInter1(t *testing.T) {
	b := new(Boy)
	b.walk()
}

func TestInter2(t *testing.T) {
	test(1)
	test("1")
}

func test(value interface{}) {
	val, ok := value.(int)
	if !ok {
		fmt.Println(ok)
	} else {
		fmt.Println(val)
	}
}
