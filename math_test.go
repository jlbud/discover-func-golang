package main

import (
	"math"
	"testing"
)

// 2的3次方
func TestMathPow(t *testing.T) {
	seed := 2
	delayTime := math.Pow(2, float64(seed))*30 + 60
	t.Log(delayTime)
}
