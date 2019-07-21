package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test_slice_int_array(t *testing.T) {
	var e []int
	e = append(e, 1)
	e = append(e, 2)
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
