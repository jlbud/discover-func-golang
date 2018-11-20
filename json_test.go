package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

//struct to json
func Test_struct_json(t *testing.T) {
	type ColorGroup struct {
		ID     int      `json:"id,string"`
		Name   string   `json:"name,omitempty"` //有omitempty，不填写该字段，则不会输出
		Colors []string `json:"colors"`         //没有omitempty，不填写该字段，则不会输出
	}
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	b, err := json.Marshal(group)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
	// output: {"id":"1","name":"Reds","colors":["Crimson","Red","Ruby","Maroon"]}

	// 如果没有设置Name属性值，因为标记为了omitempty属性，则在编码成json的时候会忽略Name属性
	group = ColorGroup{
		ID:     1,
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	b, err = json.Marshal(group)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
	// output: {"id":"1","colors":["Crimson","Red","Ruby","Maroon"]}

	// 如果没有设置Colors值，因为没有omitempty属性，会输出nil
	group = ColorGroup{
		ID:   1,
		Name: "Reds",
	}
	b, err = json.Marshal(group)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
	// output: {"id":"1","name":"Reds","colors":null}
}

//map to json
func Test_map_to_json(t *testing.T) {
	m := map[string]interface{}{
		"id":      1,
		"name":    "xiao ming",
		"friends": []string{"xiao fang", "xiao huang"},
	}
	b, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))

}
