package main

import (
	"fmt"
	"testing"
)

// go test -v -fuzz=Fuzz -fuzztime 10s -run FuzzEqual // auto input fuzz params,
// has " matches more than one fuzz test" bug.
// go test -v -run FuzzEqual // only input two where Add func() special cases.
func FuzzEqual(f *testing.F) {
	f.Add([]byte{1, 2, 3}, []byte{1, 2, 3})
	f.Add([]byte{2, 3, 4}, []byte{1, 2, 3})
	f.Fuzz(func(t *testing.T, a []byte, b []byte) {
		bol := equalByte(a, b)
		if bol {
			fmt.Println("equal")
		}else{
			fmt.Println("not equal")
		}
	})
}
func equalByte(a []byte, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
