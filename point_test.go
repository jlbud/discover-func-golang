package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// 声明指针
// 解指针
func TestPoi1(t *testing.T) {
	var a int = 4
	var pri *int      // 指针int变量
	pri = &a          // a的内存地址给pri
	fmt.Println(pri)  // 输出a的内存地址 0xc000110338
	fmt.Println(*pri) // *为解指针pri，输出a的值 4
}

// change *int for struct
type In struct {
	I int
}

func TestPointer2(t *testing.T) {
	//var in *int
	a := 1
	in := &In{
		I: a,
	}
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			in.I = i
		}
	}()
	sy := &sync.WaitGroup{}
	sy.Add(1)
	s(in, sy)
	sy.Wait()
	time.Sleep(2 * time.Second)
}
func s(p *In, sy *sync.WaitGroup) {
	defer sy.Done()
	for i := 0; i < 15; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(p.I)
	}
}
