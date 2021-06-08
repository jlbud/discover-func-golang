package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
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

func TestFloatFixedZeroDigit(t *testing.T) {
	f := 2.500000
	fixed, _ := strconv.ParseFloat(fmt.Sprintf("%.1f", f), 64)
	t.Log(fixed)
}

// 整数转浮点数相除
func TestFloat1(t *testing.T) {
	rate := float64(2) / float64(3)
	t.Log(rate)
}

func TestFloat2(t *testing.T) {

	t.Log(FixedThreeDigit(float64(8) / float64(9)))
}

func round(input float64, places int) float64 {
	coefficient := 1.0
	if input < 0.0 {
		coefficient = -1.0
		input *= -1.0
	}
	return coefficient * precision(precision(input+0.01, places+1, true), places, true)
}

// 2.25 1
func precision(f float64, prec int, round bool) float64 {
	pow10_n := math.Pow10(prec)
	if round {
		fmt.Println("pow10_n", pow10_n)
		a := f + 0.5/pow10_n // 2.2505
		fmt.Println("a", a)
		b := a * pow10_n // 2.2505 * 1000 = 2250.5
		fmt.Println("b", b)

		c := math.Trunc(b) // 2250
		fmt.Println("c", c)
		return c / pow10_n // 2.25
	}
	return math.Trunc((f)*pow10_n) / pow10_n
}

func FixedThreeDigit(f float64) float64 {
	fixed, _ := strconv.ParseFloat(fmt.Sprintf("%.1f", f), 64)
	return fixed
}

func TestOneDigit(t *testing.T) {
	//f := float64(1) / float64(9)
	//f := float64(2.255)
	f := float64(22) / float64(6)
	t.Log(decimal(f))
}

func decimal(f float64) float64 {
	s := fmt.Sprintf("%v", f)
	if !strings.Contains(s, ".") {
		return FixedThreeDigit(f)
	} else {
		//if len(strings.Split(s, ".")[1]) > 2 {
		//	fmt.Println("2222")
		//	return round(f, 2)
		//} else {
		fmt.Println("1111")
		return round(f, 1)
		//}
	}
}

func TestFloatDefault(t *testing.T) {
	var f float32
	if f == 0.0 {
		t.Logf("the float32 default value is %v", f)
	} else {
		t.Logf("the float32 value is %v", f)
	}
}
