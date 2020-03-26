package main

import (
	"fmt"
	"testing"
)

var i int

func TestReturnA(t *testing.T) {
	i = 0
	handle()
	fmt.Println("end")
}
func handle() {
	i++
	fmt.Println(i)
	if i < 10 {
	handle()
	fmt.Printf("第%d层", i)
	fmt.Println()
	}
	fmt.Println("this is ", i)
	return
}

func TestReturnB(t *testing.T) {
	type we struct {
		a string
	}
	wl := make([]we, 0)
	wl = append(wl, we{a: "we are family"})

	fmt.Printf("[AutoHome] companyID %d, batch staff len is %d", 1, len(wl))
	fmt.Println()
	fmt.Printf("[AutoHome] companyID %d, batch department len is %d", 1, len(wl))
	fmt.Println()
	fmt.Printf("[AutoHome] companyID %d, batch staff jurisdiction len is %d", 1, len(wl))
}
