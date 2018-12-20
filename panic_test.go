package main

import (
	"fmt"
	"testing"
	"time"
)

// golang panic捕捉顺序，注意当前go程
func Test_panic_order(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Error(fmt.Errorf("err log: %v", err))
		}
	}()

	go one(t)
	t.Log("start two")
	time.Sleep(5 * time.Second)
}

func one(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Error(fmt.Errorf("one inner err log: %v", err))
		}
	}()
	two()
}

func two() {
	panic("two panic")
}

// defer顺序，先输出1后输出2
func Test_defer_order(t *testing.T) {
	defer func() {
		fmt.Println("2")

	}()
	defer func() {
		fmt.Println("1")
	}()
}
