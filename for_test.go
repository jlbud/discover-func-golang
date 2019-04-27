package main

import (
	"fmt"
	"testing"
	"time"
)

/**
用break方式跳出for select循环
 */
func TestForBreak(t *testing.T) {
	i := 0
LOOP:
	for {
		select {
		case <-time.After(time.Second * time.Duration(2)):
			i ++
			if i == 5 {
				fmt.Println("跳出for循环")
				break LOOP
			}
		}
		fmt.Println("for循环内 i=", i)
	}
	fmt.Println("for循环外")
}

/**
用goto方式跳出for select循环
 */
func TestForGoto(t *testing.T) {
	i := 0
	for {
		select {
		case <-time.After(time.Second * time.Duration(2)):
			i++
			if i == 5 {
				fmt.Println("跳出for循环")
				goto LOOP
			}
		}
		fmt.Println("for循环内 i=", i)
	}
LOOP:
	fmt.Println("for循环外")
}
