package main

import (
	"fmt"
	"testing"
)

//定义接口
type adder interface {
	add(string) int
}

//定义函数类型
type handler func(name string) int

//实现函数类型方法
func (h handler) add(name string) int {
	return h(name) + 10
}

//函数参数类型接受实现了adder接口的对象（函数或结构体）
func process(a adder) {
	fmt.Println("process:", a.add("taozs"))
}

//另一个函数定义
func doubler(name string) int {
	return len(name) * 2
}

//非函数类型
type myint int

//实现了adder接口
func (i myint) add(name string) int {
	return len(name) + int(i)
}

var my handler = func(name string) int {
	return len(name)
}

func Test_func(t *testing.T) {
	//以下是函数或函数方法的调用
	fmt.Println(my("taozs"))     //调用函数
	fmt.Println(my.add("taozs")) //调用函数对象的方法
	//doubler函数显式转换成handler函数对象然后调用对象的add方法，先执行doubler函数结果为10，然后转为handler函数变量，最后执行
	//handler的add函数10 + 10 = 20
	fmt.Println(handler(doubler).add("taozs"))
	//以下是针对接口adder的调用
	process(my)               //process函数需要adder接口类型参数
	process(handler(doubler)) //因为process接受的参数类型是handler，所以这儿要强制转换
	process(myint(8))         //实现adder接口不仅可以是函数也可以是结构体
}
