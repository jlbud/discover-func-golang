package main

import (
	"fmt"
	"testing"
	"time"
)

//================================== Test_chan_equilibrium start
//用channel实现的负载均衡
func Test_chan_equilibrium(t *testing.T) {
	ch := make(chan interface{})
	go ch1(ch)
	go ch2(ch)
	for i := 0; i < 100; i++ {
		ch <- i
	}
	time.Sleep(10 * time.Second)
}
func ch1(ch chan interface{}) {
	for {
		select {
		case v, ok := <-ch:
			if ok {
				fmt.Println("a=====", v)
			}
		}
	}
}
func ch2(ch chan interface{}) {
	for {
		select {
		case v, ok := <-ch:
			if ok {
				fmt.Println("b=====", v)
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
