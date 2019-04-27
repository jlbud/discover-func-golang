package main

import (
	"fmt"
	"testing"
	"time"
)

/**
先打印aaa
再打印bbb
再打印aaa
此时natures <- x 会阻塞
0 1 4 9 16 ... 9801
 */
func TestChA(t *testing.T) {
	natures := make(chan int)
	squares := make(chan int)

	go func() {
		for x := 0; x < 100; x++ {
			//fmt.Print("aaa")
			natures <- x
			//fmt.Print("bbb")
		}
		close(natures)
	}()

	go func() {
		for x := range natures {
			squares <- x * x
		}

		close(squares)
	}()

	for x := range squares {
		fmt.Print(x, " ")
	}
}

/**
可以测试内存回收
2秒内真个100万个循环还没有跑完，内存开始疯涨
但是主线程已经结束掉，但是内存已经跌下来了
 */
func TestChB(t *testing.T) {
	go func() {
		for i := 0; i < 1000000; i++ {
			fmt.Println("aaaaaaaaaaaa", "id ", i)
		}
		time.Sleep(100 * time.Second)
	}()
	time.Sleep(2 * time.Second)
}

/**
死锁的测试
如果缓冲为是0，则发生死锁
如果缓冲大于0，则正常输出
 */
func TestChC(t *testing.T) {
	ch := make(chan int, 1)
	ch <- 1
	select {
	case <-ch:
		fmt.Println("this is output")
	}
	fmt.Println("end")
}
