package main

import (
	"fmt"
	"testing"
)

/////////////// slice
func TestGeneric(t *testing.T) {
	printSlice([]string{"a", "b"})
}

func printSlice[T []string](s T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

/////////////// different var
func TestGeneric2(t *testing.T) {
	r := add2(10, 20)
	fmt.Println(r)
	r1 := add2("10", "20")
	fmt.Println(r1)
	r2 := add2(1.12, 2.1)
	fmt.Println(r2)
}

func add2[T int | float64 | string](a, b T) T {
	return a + b
}

/////////////// aliases
type oint int

// int aliases
func add3[T ~int | ~float64](a, b T) T {
	return a + b
}

func TestGeneric3(t *testing.T) {
	var a oint = 1
	var b oint = 2
	r := add3(a, b)
	fmt.Println(r)
}
