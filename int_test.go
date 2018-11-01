package main

import (
	"testing"
	"math/big"
	"fmt"
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
}
