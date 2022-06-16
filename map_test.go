package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"sync"
	"testing"
)

//它主要五个方法及其功能简介
//1、Store   存 key,value
//2、LoadOrStore   取&存-具体看代码
//3、Load   取key对应的value
//4、Range   遍历所有的key,value
//5、Delete   删除key,及其value
func Test_sync_map(t *testing.T) {
	var m sync.Map

	//Store
	m.Store(1, "a")
	m.Store(2, "b")

	//LoadOrStore
	//若key不存在，则存入key和value，返回false和输入的value
	v, ok := m.LoadOrStore("11", "aaa")
	fmt.Println(ok, v) //false aaa，false表示不存在"11"字符串，返回存入的"aaa"字符串

	//若key已存在，则返回true和key对应的value，不会修改原来的value
	v, ok = m.LoadOrStore(1, "aaa")
	fmt.Println(ok, v) //true "a"，true表示已经有这个值，返回已经存入的"a"字符串

	//Load
	v, ok = m.Load(1)
	if ok {
		fmt.Println("it's an existing key,value is ", v)
	} else {
		fmt.Println("it's an unknown key")
	}

	//Range
	//遍历sync.Map, 要求输入一个func作为参数
	f := func(k, v interface{}) bool {
		//这个函数的入参、出参的类型都已经固定，不能修改
		//可以在函数体内编写自己的代码，调用map中的k,v
		fmt.Println(k, v)
		//true表示继续遍历，false表示结束遍历
		return true
	}
	m.Range(f)

	//Delete
	m.Delete(1)
	//返回<nil> false，表示不存在
	fmt.Println(m.Load(1))

	//Range
	m.Range(f)
}

func Test_basic_map(t *testing.T) {
	///////// 先声明map
	var m1 map[string]string
	// 再使用make函数创建一个非nil的map，nil map不能赋值
	m1 = make(map[string]string)
	// 最后给已声明的map赋值
	m1["a"] = "aa"
	m1["b"] = "bb"

	///////// 直接创建
	m2 := make(map[string]string)
	// 然后赋值
	m2["a"] = "aa"
	m2["b"] = "bb"
	///////// 初始化 + 赋值一体化
	m3 := map[string]string{
		"a": "aa",
		"b": "bb",
	}
	///////// 查找键值是否存在
	if v, ok := m3["a"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("Key Not Found")
	}

	for k, v := range m3 {
		fmt.Println(k, v)
	}

	for k := range m3 {
		fmt.Println("m3 key is ", k)
	}

	// m3的长度
	l := len(m3)
	fmt.Println("m3 length is ", l)
}

func Test_b_map(t *testing.T) {
	m1 := make(map[string]string)
	if m1 == nil {
		t.Log("m1 is nil")
	}
	var m2 map[int]string
	if m2 == nil {
		t.Log("m2 is nil")
		t.Logf("m2 len is %d", len(m2))
	}
}

type StaffCreatesRequest struct {
	Token string         `json:"token" validate:"required"`
	List  []StaffCreates `json:"list"`
}

type StaffCreates struct {
	Name           string `json:"name" validate:"required"`            // 员工姓名
	No             string `json:"no" validate:"required"`              // 员工工号
	Sex            int    `json:"sex" validate:"required"`             // 性别
	Phone          string `json:"phone"`                               // 手机号
	Email          string `json:"email" validate:"required"`           // 邮箱
	DepartmentCode string `json:"department_code" validate:"required"` // 部门编码
	PositionCode   string `json:"position_code" validate:"required"`   // 岗位编码
	Position       string `json:"position" validate:"required"`        // 岗位名字
	RealParentNo   string `json:"real_parent_no"`                      // 直属上级工号
	InventParentNo string `json:"invent_parent_no"`                    // 虚线上级工号
	OtherParentNo  string `json:"other_parent_no"`                     // 其他上级工号
	Status         int    `json:"status" validate:"required"`          // 状态
	Education      int    `json:"education"`                           // 学历
	Birthday       string `json:"birthday"`                            // 出生日期
	WorkedAt       string `json:"worked_at"`                           // 工作时间
	HiredAt        string `json:"hired_at"`                            // 入职时间
	ProductIds     string `json:"product_ids"`                         // 产品 IDs
}

type Staff struct {
	Status int           // 1: 员工号未入库, 2: 员工号已入库
	Infor  *StaffCreates //
}

var jsonS = `{
                    "token": "4f5f3b9ad2e52808cf3165d612764519",
                    "list": [
                        {
                            "name": "位位",
                            "no": "wei002",
                            "sex": 1,
                            "email": "18268343609@qq.com",
                            "phone": "18268343609",
                            "department_code": "s001",
                            "position": "ssc测试岗",
                            "position_code": "PS001",
                            "real_parent_no": "wj004",
                            "invent_parent_no": "",
                            "other_parent_no": "",
                            "status": 2,
                            "education": 8,
                            "birthday": "2020-01-01",
                            "worked_at": "2020-01-01",
                            "hired_at": "2020-01-01"
                        },
                        {
                            "name":"June",
                            "no": "wh001",
                            "sex": 2,
                            "email": "2495873219@qq.com",
                            "phone": "18268343601",
                            "department_code": "d001",
                            "position": "dhr测试岗",
                            "position_code": "pd001",
                            "real_parent_no": "vj001",
                            "invent_parent_no": "",
                            "other_parent_no": "",
                            "status": 2,
                            "education": 8,
                            "birthday": "2020-01-01",
                            "worked_at": "2020-01-01",
                            "hired_at": "2020-01-01"
                        },
                        {
                            "name": "位娟",
                            "no": "wj004",
                            "sex": 1,
                            "email": "juan.wei@ifchange.com",
                            "phone": "18268343608",
                            "department_code": "test001",
                            "position": "测试岗位",
                            "position_code": "P01",
                            "real_parent_no": "172",
                            "invent_parent_no": "",
                            "other_parent_no": "",
                            "status": 2,
                            "education": 7,
                            "birthday": "2020-01-01",
                            "worked_at": "2020-01-01",
                            "hired_at": "2020-01-01"
                        },
                        {
                            "name": "vickie",
                            "no": "vj001",
                            "sex": 1,
                            "email": "18268343608@163.com",
                            "phone": "",
                            "department_code": "d001",
                            "position": "新增岗位",
                            "position_code": "PS001",
                            "real_parent_no": "712",
                            "invent_parent_no": "",
                            "other_parent_no": "",
                            "status": 2,
                            "education": 7,
                            "birthday": "2020-01-01",
                            "worked_at": "2020-01-01",
                            "hired_at": "2020-01-01"
                        }
                    ]
                }`

func Test_c_map(t *testing.T) {
	params := &StaffCreatesRequest{}
	var (
		staffsMap = make(map[string]Staff, 0)
	)
	_ = json.Unmarshal([]byte(jsonS), &params)

	for index := range params.List {
		staffsMap[params.List[index].No] = Staff{
			Status: 1,
			Infor:  &(params.List[index]),
		}
	}

	for staffNo := range staffsMap {
		b, _ := json.Marshal(staffsMap[staffNo].Infor)
		fmt.Printf("Batch 1, staffNo: %s,staffsMap[staffNo].Infor: %s", staffNo, string(b)) // TODO rm
		fmt.Println()
		fmt.Println()
	}
	abcd(staffsMap)
	fmt.Println("================")
	for staffNo := range staffsMap {
		b, _ := json.Marshal(staffsMap[staffNo].Infor)
		fmt.Printf("Batch 1, staffNo: %s,staffsMap[staffNo].Infor: %s", staffNo, string(b)) // TODO rm
		fmt.Println()
		fmt.Println()
	}
}

func abcd(m map[string]Staff) {
	m["vj001"] = Staff{
		Status: 2,
		Infor:  m["vj001"].Infor,
	}
}

func TestIsExists(t *testing.T) {
	maping := map[string]string{"1": "2"}
	if val, ok := maping["a"]; ok {
		t.Log(val)
	} else {
		t.Error("key not found")
	}
}

func TestMapMemory(t *testing.T) {
	var m map[int]int
	mdMap(m)
	fmt.Println(m) // the map is nil
}
func mdMap(m map[int]int) {
	fmt.Println(m) // the map is nil
	m = make(map[int]int)
	m[1] = 100
	m[2] = 200
	fmt.Println(m) // point new memory location so has new map
}

var t1 = `{
    "lines": [
        {
            "end": 1.041,
            "fluency": 0.11,
            "integrity": 9,
            "pronunciation": 19.22,
            "score": 78.99,
            "speeding": 2,
            "standardScore": 99.1,
            "words": [
                {
                    "score": 3.917,
                    "subwords": [
                        {
                            "score": 50.0
                        }
                    ]
                },
                {
                    "score": 1.22,
                    "subwords": [
                        {
                            "begin": 0.411,
                            "score": 60.0
                        },
                        {
                            "score": 60.0
                        }
                    ]
                }
            ]
        }
    ]
}`

func TestMapResetZero(t *testing.T) {
	b, err := resetZeroByMode("G", t1)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(b))
}
func resetZeroByMode(mode, s string) (string, error) {
	switch mode {
	case "A", "B", "C", "D", "E", "H", "G":
		return resetLinesWords(mode, s)
	default:
		return "", errors.New("resetZero mode not match")
	}
}
func resetLinesWords(mode string, s string) (string, error) {
	var result map[string]interface{}
	err := json.Unmarshal([]byte(s), &result)
	if err != nil {
		return "", err
	}
	// return if not found field 'lines'
	if _, ok := result["lines"]; !ok {
		return s, nil
	}
	lines := result["lines"].([]interface{})
	for _, line := range lines {
		l := line.(map[string]interface{})
		if _, ok := l["fluency"]; ok {
			l["fluency"] = 0
		}
		if _, ok := l["integrity"]; ok {
			l["integrity"] = 0
		}
		if _, ok := l["pronunciation"]; ok {
			l["pronunciation"] = 0
		}
		if _, ok := l["score"]; ok {
			l["score"] = 0
		}
		if _, ok := l["speeding"]; ok {
			l["speeding"] = 1
		}
		if _, ok := l["standardScore"]; ok {
			l["standardScore"] = 0
		}
		// reset other fields in this place
		// ...
		// continue if not found field 'words'
		if _, ok := l["words"]; !ok {
			continue
		}
		words := l["words"].([]interface{})
		for _, word := range words {
			w := word.(map[string]interface{})
			if _, ok := w["score"]; ok {
				w["score"] = 0
			}
			if mode == "E" || mode == "H" || mode == "G" {
				// reset other fields in this place
				// ...
				//continue if not found field 'subwords'
				if _, ok := w["subwords"]; !ok {
					continue
				}
				subwords := w["subwords"].([]interface{})
				for _, subword := range subwords {
					sw := subword.(map[string]interface{})
					if _, ok := sw["score"]; ok {
						sw["score"] = 0
					}
				}
			}
		}
	}
	return resultMar(result)
}
func resultMar(result map[string]interface{}) (string, error) {
	b, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
