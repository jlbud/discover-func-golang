package main

import "testing"

type A struct {
	b string
}

func (a *A) a() string {
	return "a"
}

func Test_new_a(t *testing.T) {
	// 如果A 不需要初始化内部变量，用这种方式更优秀
	str := new(A).a()
	t.Log(str)
}

type B struct {
	b string
}

func (b *B) a() string {
	return b.b
}

func Test_new_b(t *testing.T) {
	// 如果B 需要初始化内部变量，用这种方式更优秀
	b := &B{
		b: "b",
	}
	t.Log(b.a())
}
