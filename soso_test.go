package main

import (
	"fmt"
	"sync/atomic"
	"testing"
)

var i2 int32 = 1
func Test_main(t *testing.T) {
	var i1 int32 = 1


	atomic.AddInt32(&i1, 3)

	fmt.Println(&i1)
	if bol := atomic.CompareAndSwapInt32(&i1, i1, i2); bol {
		fmt.Println(&i1)
	}
}
