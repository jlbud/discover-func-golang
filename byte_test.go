package main

import (
	"testing"
	"unsafe"
)

//非安全指针修改[]byte
func Test_byte_unsafe_pointer(t *testing.T) {
	bytes := []byte("i is byte")
	str := (*string)(unsafe.Pointer(&bytes))
	bytes[0] = 'y'
	t.Log(*str)
}