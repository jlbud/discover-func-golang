package main

import (
	"fmt"
	"testing"
)

//================================== Test_chan_equilibrium start
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

//================================== Test_chan_equilibrium end

//================================== Test_returnOnlyReadChan start
//返回一个只读channel
func Test_returnOnlyReadChan(t *testing.T) {
	r := returnOnlyReadChan()
	fmt.Println(<-r)
}
func returnOnlyReadChan() <-chan int {
	i := make(chan int, 1)
	i <- 1
	return i
}

//================================== Test_returnOnlyReadChan end
