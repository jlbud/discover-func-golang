package main

import (
	"fmt"
	"strconv"
	"testing"
)

func Test_float(t *testing.T) {
	var a float64
	var b float64

	a = 0
	b = 0.0
	if a == b {
		t.Log("equal")
		return
	}
	t.Log("return")
}

// float64精度
func Test_float_a(t *testing.T) {
	t.Log(float64(51) / float64(91))
	t.Log(float32(Float64Round(float64(51)/float64(91), 2)))
	t.Log(strconv.ParseFloat(fmt.Sprintf("%.2f", float64(51)/float64(91)), 64))
}

// float64精度处理
func Float64Round(f float64, n int) float64 {
	format := "%." + strconv.Itoa(n) + "f"
	res, _ := strconv.ParseFloat(fmt.Sprintf(format, f), 64)
	return res
}
