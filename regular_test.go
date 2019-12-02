package main

import (
	"regexp"
	"testing"
)

// `^1`第一位必须为1
// `[3456789]`第二位是3 4 5 6 7 8 9中一位
// `\d{9}`后边跟数字并且连续匹配9次
var regPhone = `^1[3456789]\d{9}$`

func TestRegularPhone(t *testing.T) {
	reg := regexp.MustCompile(regPhone)
	bool := reg.MatchString("13000000001")
	t.Log(bool)
}
