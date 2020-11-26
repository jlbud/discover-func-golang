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

type Friend struct {
	Friends []string `json:"friends"`
	ID      int64    `json:"id"`
	Name    string   `json:"name"`
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

type Topic struct {
	Topics []string `json:"topics"`
}

func Test_a(t *testing.T) {
	var str string
	if str == "" {
		fmt.Println("a is \"\" ")
	}
}

type CustomText struct {
	ID      int64           `json:"id"`
	Text    string          `json:"text"`
	Options []*CustomOption `json:"options"`
}

type CustomOption struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Desc  string `json:"desc"`
}

func Test_aits(t *testing.T) {
	e := &CustomText{}
	e.ID = -1
	e.Text = "关于公司文化，你希望你所在的公司是以下的哪一种呢？"

	var customs []*CustomOption

	c1 := &CustomOption{}
	c1.Key = "KeyA"
	c1.Value = "功能型"
	c1.Desc = "功能型工作文化是指以制度为核心，注重组织中上下级关系，组织的管理者权威性和员工的专业技能的工作文化。"
	customs = append(customs, c1)

	c2 := &CustomOption{}
	c2.Key = "KeyB"
	c2.Value = "流程型"
	c2.Desc = "流程型工作文化是指公司必须检视各项零散的任务，并整合为一完善的工作流程，再利用团队合作来执行整个工作流程，而达成公司目标。"
	customs = append(customs, c2)

	c3 := &CustomOption{}
	c3.Key = "KeyC"
	c3.Value = "即时型"
	c3.Desc = "即时型工作文化是指企业先将工作结构设计成专案式的工作型态，再以多技能的员工来执行任务，利用专案成员的专业知识帮助企业不断的创新。"
	customs = append(customs, c3)

	c4 := &CustomOption{}
	c4.Key = "KeyD"
	c4.Value = "网路型"
	c4.Desc = "网路型工作文化是指由于环境不确定性的增加，企业开始采用策略联盟的方式来降低本身固定成本的投入与风险的承担，并使企业结构更具变动性，透过双方合作来完成特定的任务的工作方式。"
	customs = append(customs, c4)

	e.Options = customs

	s, err := json.Marshal(e)
	if err != nil {
		return
	}
	fmt.Println(string(s))
}

func TestAJson(t *testing.T) {
	type Empty struct {
		AA int `json:"a"`
		BB int `json:"b"`
	}

	e := &Empty{
		AA: 1,
		BB: 1,
	}

	a, _ := json.Marshal(e)
	//json.Unmarshal(a,)
	fmt.Println(string(a))
}

func TestBJson(t *testing.T) {

}
