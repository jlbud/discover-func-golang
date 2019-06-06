package main

import (
	"fmt"
	"testing"
)

////////////////////////////////////////// start func接受参数的类型
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

func Test_func_parameter(t *testing.T) {
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

////////////////////////////////////////// end func接受参数的类型

////////////////////////////////////////// start func作为参数转换
func Test_func_conversion(t *testing.T) {
	//将language函数转为Man类型
	m := Man(language)
	//执行Man类型的一个方法
	m.action("hello")

	//将run函数转为Man类型
	r := Man(run)
	//执行Man类型的一个方法
	r.action("10")
}

//声明一个函数类型
type Man func(name string) string

func (m Man) action(msg string) {
	s := m(msg) //使用函数
	fmt.Println(s)
}

//定义一个函数，格式和type Man相同
func language(msg string) string {
	return "say " + msg
}

//定义一个函数，格式和type Man相同
func run(speed string) string {
	return "speed " + speed
}

////////////////////////////////////////// end func作为参数转换

////////////////////////////////////////// start func作为回调参数
func Test_func_type_callback(t *testing.T) {
	//调用callback函数
	i := callback(1, 2, addCallBack)
	fmt.Println(i)
}

func addCallBack(a, b int) int {
	return a + b
}

type Callback func(int, int) int

func callback(x, y int, cb Callback) int {
	i := cb(x, y)
	return i
}

////////////////////////////////////////// end func作为回调参数

func Test_func_callback(t *testing.T) {
	from(to)
}

func to(s string) {
	fmt.Println(s)
}

func from(to func(s string)) {
	to("hello")
}
