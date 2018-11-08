package main

import (
	"testing"
	"fmt"
)

func Test_main(t *testing.T) {
	a := newInt()
	str := "a"
	a.Read = &str
	fmt.Println(a.Read)
}

type ios struct {
	Read *string
}

func newInt() *ios {
	var i ios

	return &i
}
