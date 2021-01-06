package main

import (
	"fmt"
	"testing"
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
