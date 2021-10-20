package main

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
)

type data struct {
	num    int
	checks [10]func() bool   //无法比较
	doIt   func() bool       //无法比较
	m      map[string]string //无法比较
	bytes  []byte            //无法比较
}

//比较两个结构体是否相等，因为有无法比较的字段在
func Test_deep_equal(t *testing.T) {
	v1 := data{}
	v2 := data{}
	fmt.Println("v1 == v2: ", reflect.DeepEqual(v1, v2)) // true

	m1 := map[string]string{"one": "a", "two": "b"}
	m2 := map[string]string{"two": "b", "one": "a"}
	fmt.Println("v1 == v2: ", reflect.DeepEqual(m1, m2)) // true

	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	// 注意两个 slice 相等，值和顺序必须一致
	fmt.Println("v1 == v2: ", reflect.DeepEqual(s1, s2)) // true
}

//字节比较
func Test_bytes_equsl(t *testing.T) {
	var b1 []byte = nil
	b2 := []byte{}
	// b1 与 b2 长度相等、有相同的字节序
	// nil 与 slice 在字节上是相同的
	fmt.Println("b1 == b2: ", bytes.Equal(b1, b2)) // true
}
