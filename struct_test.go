package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

// type，有些时候，我们在序列化或者反序列化的时候，可能结构体类型和需要的类型不一致，这个时候可以指定,支持string,number和boolean

// Product
type Product struct {
	ProductID int64 `json:"product_id,string"`
	Num       bool  `json:"num,string"`
}

func TestS1(t *testing.T) {
	//var data = `{"name":"Xiao mi 6","product_id":"10","num":10000,"price":"2499","is_on_sale":"true"}`
	var data = `{"product_id":"10","num":"100"}`
	p := &Product{}
	_ = json.Unmarshal([]byte(data), p)
	fmt.Println(*p)
}

func TestStructValidator(t *testing.T) {
	v := struct {
		A string `must:"true"`
		B string `must:"false"`
	}{
		A: "a",
	}
	err := validator(&v) // must be input pointer
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	t.Log("success")
}

func validator(i interface{}) error {
	val := reflect.ValueOf(i).Elem() // get all fields value
	ty := reflect.TypeOf(i).Elem()   // get all fields type
	for i := 0; i < ty.NumField(); i++ {
		if ty.Field(i).Tag.Get("must") == "true" && val.Field(i).IsZero() {
			return fmt.Errorf("<%v> not be empty", ty.Field(i).Name)
		}
	}
	return nil
}
