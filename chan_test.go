package main

import (
	"fmt"
	"testing"
)

//用channel实现的负载均衡
func Test_chan_equilibrium(t *testing.T) {
	ch := make(chan interface{})
	go a(ch)
	go b(ch)
	for i := 0; i < 100; i++ {
		ch <- i
	}
}
func a(ch chan interface{}) {
	count := 0
	for {
		select {
		case _, ok := <-ch:
			if ok {
				count ++
				fmt.Println("a=====", count)
			}
		}
	}
}
func b(ch chan interface{}) {
	count := 0
	for {
		select {
		case _, ok := <-ch:
			if ok {
				count ++
				fmt.Println("b=====", count)
			}
		}
	}
}
