package weird

import (
	"fmt"
	"testing"
)

//给slice赋值nil
func Test_slice_nil(t *testing.T) {
	s := make([]string, 1)
	s[0] = "0"
	sliceNil(s)

	fmt.Println()
	fmt.Printf("内存地址:%p \t\t长度:%v\t\t容量:%v\t\t包含元素:%v\n", s, len(s), cap(s), s)
}

func sliceNil(s []string) {
	s[0] = "1"
	fmt.Printf("内存地址:%p \t\t长度:%v\t\t容量:%v\t\t包含元素:%v\n", s, len(s), cap(s), s)

	fmt.Println()
	//赋值nil会让s重新指向新的内存地址，只能在内部使用，外部的s内存地址不会改变
	s = nil
	//重新给新的s分配内存
	s = make([]string, 1)
	s[0] = "1000"
	fmt.Printf("内存地址:%p \t\t长度:%v\t\t容量:%v\t\t包含元素:%v\n", s, len(s), cap(s), s)
}
