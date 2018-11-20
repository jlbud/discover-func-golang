package main

import (
	"fmt"
	"testing"
)

func Test_main(t *testing.T) {
	s := make([]string, 1)
	s[0] = "0"
	a(s)

	fmt.Println()
	fmt.Printf("内存地址:%p \t\t长度:%v\t\t容量:%v\t\t包含元素:%v\n", s, len(s), cap(s), s)
}

func a(s []string) {
	s[0] = "1"

	fmt.Printf("内存地址:%p \t\t长度:%v\t\t容量:%v\t\t包含元素:%v\n", s, len(s), cap(s), s)

	fmt.Println()
	s = nil
	s = make([]string, 1)
	s[0] = "1000"
	fmt.Printf("内存地址:%p \t\t长度:%v\t\t容量:%v\t\t包含元素:%v\n", s, len(s), cap(s), s)
}
