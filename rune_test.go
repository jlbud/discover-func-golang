package main

import (
	"fmt"
	"testing"
	"unicode"
	"unicode/utf8"
)

func TestRune1(t *testing.T) {
	s := "白日依上进"
	for _, r := range s {
		fmt.Println("r===>", r)
		fmt.Printf("%c = %v\n", r, unicode.IsSpace(r))
	}
}

func TestRune2(t *testing.T) {
	s := "世界我们"
	_, size := utf8.DecodeRuneInString(s)
	t.Log("size", size)
	t.Log("s[:size]", s[:size])
	t.Log([]string{s[:size]})
}

func FindHan(s string) []int {
	var loc [2]int

	_, size := utf8.DecodeRuneInString(s)
	fmt.Println("size", size)
	if size <= 1 {
		fmt.Println("size111")
		loc[1] = size
		return loc[:]
	}

	loc[1] = size
	return loc[:]
}

func TestRune3(t *testing.T) {
	str := "Hello, 世界"

	for len(str) > 0 {
		r, size := utf8.DecodeRuneInString(str)
		fmt.Printf("%c %v\n", r, size)

		str = str[size:]
		t.Log(str)
	}
}
