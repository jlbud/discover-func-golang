package main

import (
	"context"
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
		for i := 0; i < 100; i++ {
			fmt.Println("aaaaaaaaaaaa", "id ", i)
		}
		time.Sleep(100 * time.Second)
	}()
	time.Sleep(2 * time.Second)
}

type ABC struct {
	A int
	B int
}

// 死锁的测试
// 如果缓冲为是0，则发生死锁
// 如果缓冲大于0，则正常输出
// 如果close()，则退出
func TestChOfClose(t *testing.T) {
	ch := make(chan ABC, 10)
	ch1 := make(chan ABC, 10)

	ch <- ABC{A: 1, B: 2}
	ch <- ABC{A: 4, B: 3}
	ch <- ABC{A: 2, B: 8}
	ch <- ABC{A: 2, B: 8}
	close(ch)
FOR:
	for {
		select {
		case i, ok := <-ch:
			fmt.Println(ok)
			if ok {
				fmt.Println("this is output:", i.A, i.B, ok)
			} else {
				break FOR
			}
		case _, ok := <-ch1:
			fmt.Println(ok)
		}
	}
	fmt.Println("end")
}

func TestChD(t *testing.T) {
	ch := make(chan int, 3)
	ch <- 1
	l := len(ch)
	c := cap(ch)
	fmt.Println(l)
	fmt.Println(c)
}

func TestCh1(t *testing.T) {
	in := make(chan int)
	go testCh1(in)
	select {
	case v := <-in:
		t.Log(v)
	case <-time.After(time.Second * 2):
		t.Log("timeout...")
	}
	t.Log("logout...")
}

func testCh1(in chan int) {
	time.Sleep(time.Second * 10)
	in <- 1
}

func TestCh2(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	select {
	case <-ctx.Done():
		t.Log("done")
	}
}
