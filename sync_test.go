package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"sync"
	"testing"
	"time"
)

var l sync.Mutex

func TestSyncMutex(t *testing.T) {
	a := 0
	p_a := &a
	go func() {
		l.Lock()
		fmt.Println("3")
		time.Sleep(5 * time.Second)
		*p_a += 1
		fmt.Println("4")
		l.Unlock()
	}()

	go func() {
		l.Lock()
		fmt.Println("1")
		*p_a += 1
		fmt.Println("2")
		l.Unlock()
	}()

	go func() {
		l.Lock()
		time.Sleep(3 * time.Second)
		fmt.Println("5")
		*p_a += 1
		fmt.Println("6")
		l.Unlock()
	}()
	fmt.Println(*p_a)

	time.Sleep(15 * time.Second)
}

