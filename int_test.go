package main

import (
	"fmt"
	"math/big"
	"strconv"
	"testing"
)

//string转大整数
func Test_string_to_bigint(t *testing.T) {
	amount := new(big.Int)
	amount.SetString("100.00", 10)
	fmt.Println(amount)
}

//int64转大整数
func Test_int64_to_bigint(t *testing.T) {
	limitdate := new(big.Int)
	limitdate.SetInt64(10)
	fmt.Println(limitdate)
}

func Test_int_to_int32(t *testing.T) {
	var i int
	i = 999999999 // max length to 32 right
	i32 := int32(i)
	t.Logf("int: %d", i)
	t.Logf("int32: %d", i32)
}

// int到string
func Test_int_to_string(t *testing.T) {
	string := strconv.Itoa(1)
	t.Log(string)
}

func Test_max_int(t *testing.T) {
	var a int
	a = 1000000000000000000 // int型最多的位数
	a = 999999999999999999  // int型最大的值
	b := a % 5
	t.Log(b)
}

func Test_int_division(t *testing.T) {
	result := int(86401 / 86400)
	t.Log(result)
}

func TestInt1(t *testing.T) {
	str1 := "2"
	str2 := "1"
	if str1 < str2 {
		t.Logf("%s<%s", str1, str2)
	} else {
		t.Logf("%s>%s", str1, str2)
	}
}
