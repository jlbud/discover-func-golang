package main

import (
	"fmt"
	"testing"
)

type Animal interface {
	Speak() string
	Run() string
}

type Dog struct{}

func (d *Dog) Speak() string {
	return "d www"
}

func (d *Dog) Run() string {
	return "d rrr"
}

type Cat struct{}

func (c Cat) Speak() string {
	return "c mmm"
}

func (c Cat) Run() string{
	return "c rrr"
}

func animal(ani Animal) {
	fmt.Println(ani.Run())
	fmt.Println(ani.Speak())
}

func Test_main(t *testing.T) {
	animal(&Dog{})

	//var ani Animal = &Dog{}
	//fmt.Println(ani.Run())
	//fmt.Println(ani.Speak())

	//var ani1 Animal = &Cat{}
	//fmt.Println(ani1.Run())
	//fmt.Println(ani1.Speak())

}
