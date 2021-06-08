package main

import (
	"bytes"
	"encoding/gob"
	"testing"
)

type CopyA struct {
	Start int
	End   int
	// this 'addition' name in 'CopyA' and 'CopyB' must be equal and
	// the 'name' first char must be Upper
	Addition
}

type Addition struct {
	Middle string
}

type CopyB struct {
	Start    int
	End      int
	Addition // this 'addition' name in 'CopyA' and 'CopyB' must be equal
}

func TestDeepCopy(t *testing.T) {
	a := &CopyA{}
	a.Start = 1
	a.End = 1
	a.Middle = "middle"

	b := &CopyB{}
	err := deepCopy(a, b)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("a address:%v", &a)
	t.Logf("b address:%v", &b)

	t.Logf("a value:%+v", a)
	t.Logf("b value:%+v", b)

	//output:
	//copy_test.go:38: a address:0xc000196510
	//copy_test.go:39: b address:0xc000196518
	//copy_test.go:41: a value:&{Start:1 End:1 Addition:{Middle:middle}}
	//copy_test.go:42: b value:&{Start:1 End:1 Addition:{Middle:middle}}
}

func deepCopy(src, dst interface{}) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}

func TestShallowCopy(t *testing.T) {
	a := &CopyA{}
	a.Start = 1
	a.End = 1
	a.Middle = "middle"
	// 浅拷贝
	lowCopyA := a
	t.Logf("a address:%v", &a)
	t.Logf("lowCopyA address:%v", &lowCopyA)

	t.Logf("a value:%+v", a)
	t.Logf("lowCopyA value:%+v", lowCopyA)

	//output:
	//copy_test.go:60: a address:0xc0000ae518
	//copy_test.go:61: lowCopyA address:0xc0000ae520
	//copy_test.go:63: a value:&{Start:1 End:1 Addition:{Middle:middle}}
	//copy_test.go:64: lowCopyA value:&{Start:1 End:1 Addition:{Middle:middle}}
}
