package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 1) //修改2为1就报错，修改2为3可以正常运行
	go func() {}()
	//c <- 1
	select {
	case i := <-c:
		fmt.Println(i)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout")
	}
}
