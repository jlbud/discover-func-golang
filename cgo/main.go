package main

/*
#include "hello.c"
*/
import "C"
import "fmt"

func main() {
	_, err := C.Hello()
	if err != nil {
		fmt.Println("err:", err)
	}
}
