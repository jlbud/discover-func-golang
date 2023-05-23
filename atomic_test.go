package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

//================================== Test_JustDoOnce start

type Once struct {
	m    sync.Mutex
	done uint32
}

func (o *Once) do(f func()) {
	//利用原子特性判断
	if atomic.LoadUint32(&o.done) == 1 {
		return
	}
	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 0 {
		//利用原子特性载值
		defer atomic.StoreUint32(&o.done, 1)
		f()
	}
}

// sync.once 利用原子特性的单例(单次执行)实现
func Test_justDoOnce(t *testing.T) {
	o := &Once{}
	go o.do(func() { fmt.Println("do once") })
	go o.do(func() { fmt.Println("do once") })
}

//================================== Test_JustDoOnce end
