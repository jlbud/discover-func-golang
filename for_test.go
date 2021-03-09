package main

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

/**
用break方式跳出for select循环
*/
func TestForBreak(t *testing.T) {
	i := 0
LOOP:
	for {
		select {
		case <-time.After(time.Second * time.Duration(2)):
			i++
			if i == 5 {
				fmt.Println("跳出for循环")
				break LOOP
			}
		}
		fmt.Println("for循环内 i=", i)
	}
	fmt.Println("for循环外")
}

/**
用goto方式跳出for select循环
*/
func TestForGoto(t *testing.T) {
	i := 0
	for {
		select {
		case <-time.After(time.Second * time.Duration(2)):
			i++
			if i == 5 {
				fmt.Println("跳出for循环")
				goto LOOP
			}
		}
		fmt.Println("for循环内 i=", i)
	}
LOOP:
	fmt.Println("for循环外")
}

// for循环内赋值
func TestForAssignment(t *testing.T) {
	// 获取到的是内存
	// 传指针
	type Item struct {
		name string
	}
	cardArray := make([]*Item, 0)
	cardArray = append(cardArray, &Item{name: "bird"})
	for _, card := range cardArray {
		card.name = "panda"
	}
	fmt.Println("cardArray[0].name is: ", cardArray[0].name)

	// slice传指针
	// 输出：
	// Item2 name is:  {no 1}
	// Item2 name is:  {no 2}
	type Item2 struct {
		Name string
	}
	data := make([]*Item2, 0)
	src := []*Item2{
		{Name: "no 1"},
		{Name: "no 2"},
	}
	for _, m := range src {
		data = append(data, m)
	}
	for _, s := range data {
		fmt.Println("Item2 name is: ", *s)
	}

	// slice传值
	// 如果是非指针slice，在执行for循环的时候，golang会首先创建一块内存，用于存放item，
	// 之后依次将list中的元素拷贝到这块内存，在for之后若没有继续引用便进行释放，所以在此过程中，
	// 修改item或将item放入其他的map中，只会放入最后一个元素。
	// 输出：
	type Item3 struct {
		Name string
	}
	data1 := make([]*Item3, 0)
	src1 := []Item3{
		{Name: "no 1"},
		{Name: "no 2"},
	}
	for _, m := range src1 {
		data1 = append(data1, &m)
	}
	for _, s := range data1 {
		fmt.Println("Item3 name is: ", *s)
	}

	// 获取到的是数值
	// 传值的拷贝
	intArray := []int{1, 2, 3}
	for _, val := range intArray {
		val++
	}
	fmt.Println("intArray is: ", intArray)

	// 获取到的是原始内存
	intArray1 := []int{1, 2, 3}
	for i := range intArray {
		intArray1[i]++
	}
	fmt.Println("intArray1 is: ", intArray1)
}

func TestFor1(t *testing.T) {
	type student struct {
		Name string
		Age  int
	}
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}

	for _, stu := range stus {
		// 错误打印
		// &stu指针最终指向的是stus最后一个，
		// 这样map中每个value的值就会是最后一个
		m[stu.Name] = &stu
		// 正确打印
		//m[stu.Name] = &stus[i]
		// 打印&stus[i]指针
		t.Logf("%p", &stu)
		t.Log(&stu)
	}
	for key, value := range m {
		t.Logf("key=%v,value=%v", key, value)
	}
}

func TestFor2(t *testing.T) {
	pls := [][]string{
		{"C", "C++"},
		{"JavaScript"},
		{"Go", "Rust"},
	}
	for _, v1 := range pls {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Println()
	}
}

/**
for循环，用goroutine打印数组
*/
func TestFor3(t *testing.T) {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("A: ", i)
			wg.Done()
		}()
	}

	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("B: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
