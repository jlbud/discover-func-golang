package main

import (
	"testing"
	"os"
	"fmt"
	"io"
	"io/ioutil"
	"bufio"
)

var wireteString = "测试n"
var filename = "./output_1"
var f *os.File
var err1 error

//判断文件是否存在  存在返回 true 不存在返回false
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

//判断文件是否存在
func Test_check_file_is_exist(t *testing.T) {
	bol := checkFileIsExist("./output_1")
	t.Log(bol)
}

//第一种方式，io.WriteString 写入文件
func Test_io_writestring(t *testing.T) {
	if checkFileIsExist(filename) { //如果文件存在
		f, err1 = os.OpenFile(filename, os.O_APPEND, 0666) //打开文件
		if err1 != nil {
			t.Error(err1.Error())
			return
		}
		fmt.Println("文件存在1")
	} else {
		f, err1 = os.Create(filename) //创建文件
		if err1 != nil {
			t.Error(err1.Error())
			return
		}
		fmt.Println("文件不存在2")
	}

	n, err1 := io.WriteString(f, wireteString) //写入文件(字符串)
	if err1 != nil {
		t.Error(err1.Error())
		return
	}
	fmt.Printf("写入 %d 个字节n", n)
}

//第二种方式: 使用 ioutil.WriteFile 写入文件
func Test_ioutil_writefile(t *testing.T) {
	var d1 = []byte(wireteString)
	err2 := ioutil.WriteFile("./output2", d1, 0666) //写入文件(字节数组)
	if err2 != nil {
		t.Error(err1.Error())
		return
	}
	t.Log("ioutil write file is success")
}

//第三种方式:  使用 File(Write,WriteString) 写入文件
func Test_file_write_writestring(t *testing.T) {
	f, err3 := os.Create("./output3") //创建文件
	if err3 != nil {
		t.Error(err3)
		return
	}
	defer f.Close()
	var d1 = []byte(wireteString)
	n2, err3 := f.Write(d1) //写入文件(字节数组)
	if err3 != nil {
		t.Error(err3.Error())
		return
	}
	fmt.Printf("写入 %d 个字节n", n2)
	n3, err3 := f.WriteString("writesn") //写入文件(字节数组)
	if err3 != nil {
		t.Error(err3.Error())
		return
	}
	fmt.Printf("写入 %d 个字节n", n3)
	f.Sync()
}

//第四种方式:  使用 bufio.NewWriter 写入文件
func Test_bufio_newwriter(t *testing.T) {
	f, err4 := os.Create("./output4") //创建文件
	if err4 != nil {
		t.Error(err4)
		return
	}
	defer f.Close()

	w := bufio.NewWriter(f) //创建新的 Writer 对象
	n4, err4 := w.WriteString("bufferedn")
	if err4 != nil {
		t.Error(err4.Error())
		return
	}
	fmt.Printf("写入 %d 个字节n", n4)
	w.Flush()
	f.Close()
}
