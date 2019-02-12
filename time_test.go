package main

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeParse(t *testing.T) {
	//获取本地location
	toBeCharge := "2015-01-01 00:00:00"                             //待转化为时间戳的字符串 注意 这里的小时和分钟还要秒必须写 因为是跟着模板走的 修改模板的话也可以不写
	timeLayout := "2006-01-02 15:04:05"                             //转化所需模板，必须用这个模版据说是因为是golang诞生的日期
	loc, _ := time.LoadLocation("Local")                            //重要：获取时区
	theTime, _ := time.ParseInLocation(timeLayout, toBeCharge, loc) //使用模板在对应时区转化为time.time类型
	sr := theTime.Unix()                                            //转化为时间戳 类型是int64
	fmt.Println("theTime ", theTime)                                //打印输出theTime 2015-01-01 15:15:00 +0800 CST
	fmt.Println("sr ", sr)                                          //打印输出时间戳 1420041600

	//时间戳转日期
	dataTimeStr := time.Unix(1549940849, 0).Format(timeLayout) //设置时间戳 使用模板格式化为日期字符串
	fmt.Println("dataTimeStr ", dataTimeStr)
}
