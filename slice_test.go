package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

type Array1 struct {
	I  int
	I1 int
}

func TestSliceStructArray(t *testing.T) {
	e := make([]*Array1, 0)
	e = append(e, &Array1{I: 1})
	e = append(e, &Array1{I: 2})
	b, _ := json.Marshal(e)

	err := json.Unmarshal(b, &e)
	if err != nil {
		t.Error(err)
		return
	}

	var i int
	i = 100
	for _, v := range e {
		i++
		v.I1 = i
	}

	b, _ = json.Marshal(e)
	fmt.Println(string(b))
}

func Test_slice_int_array(t *testing.T) {
	var e []int
	e = append(e, 1)
	e = append(e, 2)
	t.Log("e:", e)

	b, _ := json.Marshal(e)

	err := json.Unmarshal(b, &e)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(e)
}

func Test_slice_obj(t *testing.T) {
	type Obj struct {
		Key int `json:"key"`
	}

	b, _ := json.Marshal(&Obj{Key: 1})

	var o Obj
	err := json.Unmarshal(b, &o)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(o.Key)
}

// slice转为逗号分割的字符串
func TestSlickToArrayString(t *testing.T) {
	IDs := make([]int, 0)
	IDs = append(IDs, 1)
	IDs = append(IDs, 2)

	// 去掉slice转为字符串后的前后中括号
	s := strings.Trim(fmt.Sprint(IDs), "][")
	// 把空格转为逗号
	s1 := strings.Replace(s, " ", ",", -1)
	t.Log(s1)
}
