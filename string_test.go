package main

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
	"unicode/utf8"
)

// 打印字符串中字符个数
func Test_string_count(t *testing.T) {
	char := "中国"
	fmt.Println(utf8.RuneCountInString(char)) // 1
}

// 判断字符串是否是 UTF8 文本
func Test_is_utf8_str(t *testing.T) {
	str1 := "ABC"
	fmt.Println(utf8.ValidString(str1)) // true
	str2 := "A\xfeC"
	fmt.Println(utf8.ValidString(str2)) // false
	str3 := "A\\xfeC"
	fmt.Println(utf8.ValidString(str3)) // true    // 把转义字符转义成字面值
}

// 更新字符串
func Test_update_str(t *testing.T) {
	x := "text"
	xRunes := []rune(x)[:4]
	xRunes[0] = '我'
	x = string(xRunes)
	fmt.Println(x) // 我ext
}

// string到int
func Test_string_to_int(t *testing.T) {
	int, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(int)
}

// string到int64
func Test_string_to_int64(t *testing.T) {
	int64, err := strconv.ParseInt("123", 10, 64)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(int64)
}

// int64到string
func Test_int64_to_string(t *testing.T) {
	int64 := int64(1234)
	string := strconv.FormatInt(int64, 10)
	fmt.Println(string)
}

// Hasprefix判断字符串s是否以prefix开头
func Test_string_first_is_prefix(t *testing.T) {
	str := "abcd"
	if strings.HasPrefix(str, "a") {
		t.Log("yes is first")
		return
	}
	t.Log("no is not first")
}

// HasSuffix判断字符串s是否以suffix结尾
func Test_string_last_is_prefix(t *testing.T) {
	str := "abcd"
	if strings.HasSuffix(str, "d") {
		t.Log("yes is last")
		return
	}
	t.Log("no is not last")
}

// Contains判断字符串s是否包含substr
func Test_string_contains(t *testing.T) {
	str := "abcd"
	if strings.Contains(str, "bc") {
		t.Log("yes is contains")
		return
	}
	fmt.Println("no is not contains")
}

// Index返回字符串str在字符串s中的索引(str的第一个字符的索引)，-1表示字符串s不包含字符串str
func Test_string_index(t *testing.T) {
	int := strings.Index("abacdef", "b")
	t.Log(int)
}

// LastIndex返回字符串str在字符串s中最后出现的位置
func Test_string_substr_index(t *testing.T) {
	int := strings.LastIndex("abfcdef", "f")
	t.Log(int)
}

// 非ASCII编码的字符定位
func Test_string_use_int_index(t *testing.T) {
	str := strings.IndexRune("NLT_abc", 'b')
	str1 := strings.IndexRune("NLT_abc", 's')
	str2 := strings.IndexRune("我是中国人", '中')
	fmt.Println(str)
	fmt.Println(str1)
	fmt.Println(str2)
}

// Replace用于将str中前n个字符串old替换为字符串new，如果为-1则替换所有字符串old为new
func Test_string_change_old_to_new(t *testing.T) {
	str := strings.Replace("abcdef", "abc", "***", 3)
	fmt.Println(str)
}

// Count用于计算str在s中出现的非重叠次数
func Test_string_repeat_count(t *testing.T) {
	str := "abca"
	coun := strings.Count(str, "a")
	fmt.Println(coun)

}

// Repeat用于重复count次字符串s并返回一个新的字符串
func Test_string_repeat(t *testing.T) {
	str := "a="
	strr := strings.Repeat(str, 2)
	fmt.Println(strr)
}

// ToLower将字符串中Unicode字符全部转换为小写字符
func Test_string_to_lower(t *testing.T) {
	str := "ABCD"
	strr := strings.ToLower(str)
	fmt.Println(strr)
}

// ToUpper将字符串中Unicode字符全部转换为大写字符
func Test_string_to_upper(t *testing.T) {
	str := "abcd"
	strr := strings.ToUpper(str)
	fmt.Println(strr)
}

// TrimSpace剔除字符串开头和结尾的空白符号
// Trim剔除开头和结尾的字符串
// TrimLeft,TrimRight剔除开头或结尾字符串
// Fields利用一个或多个空白符号来作为动态长度的分隔符将字符串分割，返回slice
func Test_string_to_slice_use_blank(t *testing.T) {
	str := "a b cd"
	ali := strings.Fields(str)
	fmt.Println(ali[0])
	fmt.Println(ali[1])
	fmt.Println(ali[2])
}

// Split用自定义分割符号来对指定字符串进行分割
func Test_string_to_slice(t *testing.T) {
	str := "a*b*c"
	sli := strings.Split(str, "*")
	fmt.Println(sli[0])
	fmt.Println(sli[1])
	fmt.Println(sli[2])
}

// Join用于将元素类型为string的slice使用分割符号来拼接组成一个字符串
func Test_slice_to_string(t *testing.T) {
	sli := make([]string, 5)
	sli[0] = "a"
	sli[1] = "b"
	sli[2] = "c"
	sli[3] = "d"
	sli[4] = "e"
	resStr := strings.Join(sli, "*")
	fmt.Println(resStr)
}

// 截取子字符串
func Test_string_cut(t *testing.T) {
	var str = `<br/>{"header":}`
	i := strings.Index(str, `{`)
	sLen := len(str)
	subStr := str[i:sLen]
	fmt.Println(subStr)
}

func Test_StringSliceEqualBCE(t *testing.T) {
	a := []string{"1"}
	b := []string{"2"}
	StringSliceEqualBCE(a, b)
}

// StringSliceEqualBCE use BCE feature to optimize StringSliceEqual
// 两个[]string是否相等
func StringSliceEqualBCE(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}
	if (s1 == nil) != (s2 == nil) {
		return false
	}
	// this line can ensure the next b[i] never out of index in for...range loop
	s2 = s2[:len(s1)]
	for i, v := range s1 {
		if v != s2[i] {
			return false
		}
	}
	return true
}

func Test_newline(t *testing.T) {
	str := "new\nline"
	fmt.Println(str)
}

func TestStringTrimSuffix(t *testing.T) {
	s := "Hello world!!!!"
	ts := strings.TrimSuffix(s, "!!!!")
	fmt.Printf("%q\n", ts)
}

func TestString1(t *testing.T) {
	descL := make([]string, 0)
	descL = append(descL, fmt.Sprintf("%s。", "1"))
	descL = append(descL, fmt.Sprintf("%s。", "2"))
	t.Log(strings.Join(descL, ""))
}

// Output 'a' ascii code
func TestString2(t *testing.T) {
	a := "a"
	b := a[len(a)-1]
	t.Log(b)
}
