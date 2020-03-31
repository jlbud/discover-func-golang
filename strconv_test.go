package main

import (
	"strconv"
	"testing"
)

func TestFormantFloat(t *testing.T) {
	str := strconv.FormatFloat(-1, 'f', -1, 32)
	t.Log(str)
}
