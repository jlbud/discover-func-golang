package main

import "testing"

// go test -fuzz=Fuzz -fuzztime 10s -run FuzzEqual
func FuzzEqual(t *testing.F) {
	t.Add([]byte{1, 2, 3}, []byte{1, 2, 3})
	t.Fuzz(func(t *testing.T, a []byte, b []byte) {
		bol := equalByte(a, b)
		if bol {
			t.Log("equal")
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
