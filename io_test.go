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

//第四种方式:  使用 bufio.NewWriter 写入文件，有缓存的 文件流的读写速度要快很多
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

// fileName:文件名字(带全路径)
// content: 写入的内容
func Test_append_to_file(t *testing.T) {
	fileName := "output"
	content := "this is out content"
	// 以只写的模式，打开文件
	//O_RDONLY：只读模式(read-only)
	//O_WRONLY：只写模式(write-only)
	//O_RDWR：读写模式(read-write)
	//O_APPEND：追加模式(append)
	//O_CREATE：文件不存在就创建(create a new file if none exists.)
	//O_EXCL：与 O_CREATE 一起用，构成一个新建文件的功能，它要求文件必须不存在(used with O_CREATE, file must not exist)
	//O_SYNC：同步方式打开，即不使用缓存，直接写入硬盘
	//O_TRUNC：打开并清空文件
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0664)
	if err != nil {
		fmt.Println("cacheFileList.yml file create failed. err: " + err.Error())
	} else {
		// 查找文件末尾的偏移量
		n, _ := f.Seek(0, os.SEEK_END)
		// 从末尾的偏移量开始写入内容
		_, err = f.WriteAt([]byte(content), n)
		t.Log("write successful")
	}
	defer f.Close() // this is tag v1
}
