package main

import (
	"testing"
	"fmt"
	"unicode/utf8"
)

//打印字符串中字符个数
func Test_string_count(t *testing.T) {
	char := "中国"
	fmt.Println(utf8.RuneCountInString(char)) // 1
}

//判断字符串是否是 UTF8 文本
func Test_is_utf8_str(t *testing.T) {
	str1 := "ABC"
	fmt.Println(utf8.ValidString(str1)) // true
	str2 := "A\xfeC"
	fmt.Println(utf8.ValidString(str2)) // false
	str3 := "A\\xfeC"
	fmt.Println(utf8.ValidString(str3)) // true    // 把转义字符转义成字面值
}

//更新字符串
func Test_update_str(t *testing.T) {
	x := "text"
	xRunes := []rune(x)
	xRunes[0] = '我'
	x = string(xRunes)
	fmt.Println(x) // 我ext
}
