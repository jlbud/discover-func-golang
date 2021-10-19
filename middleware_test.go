package main

import (
	"fmt"
	"testing"
	"time"
)

func work(name string) {
	fmt.Println("hello ", name)
	time.Sleep(time.Second)
}

func calTime(f func(name string)) func(string) {
	t := time.Now()
	return func(n string) {
		defer func() {
			fmt.Println("time is ", time.Since(t))
		}()
		f(n)
	}
}

// go test -v -run TestMiddle
func TestMiddle(t *testing.T) {
	// this step just init before return func
	s := calTime(work)
	// emission the func up step
	s("world")
}
