package main

import (
	"fmt"
	"strconv"
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

func TestCompare(t *testing.T) {

	//返回现在时间
	tNow := time.Now()
	//时间转化为string，layout必须为 "2006-01-02 15:04:05"
	timeNow := tNow.Format("2006-01-02 15:04:05")
	fmt.Println("tNow(time format):", tNow)
	fmt.Println("tNow(string format):", timeNow)

	//string转化为时间，layout必须为 "2006-01-02 15:04:05"
	time, _ := time.Parse("2006-01-02 15:04:05", "2014-06-15 08:37:18")
	fmt.Println("t(time format)", time)

	//某个时间点 前后判断
	trueOrFalse := time.After(tNow)
	if trueOrFalse == true {
		fmt.Println("t（2014-06-15 08:37:18）在tNow之后!")
	} else {
		fmt.Println("t（2014-06-15 08:37:18）在tNow之前!")
	}
	fmt.Println()
}

func TestTimeNow(t *testing.T) {
	n := time.Now().Format("2006-01-02 15:04:05")
	t.Log(n)
}

// 多时区
func TestMultipleTimeZones(t *testing.T) {
	now := time.Now()
	local1, err1 := time.LoadLocation("") // 等同于"UTC"
	if err1 != nil {
		fmt.Println(err1)
	}
	local2, err2 := time.LoadLocation("Local") // 服务器设置的时区
	if err2 != nil {
		fmt.Println(err2)
	}
	local3, err3 := time.LoadLocation("America/Los_Angeles") // PDT
	if err3 != nil {
		fmt.Println(err3)
	}
	// 向东每个时区+1，向西每个时区-1
	fmt.Println(now.In(local1)) // UTC(Universal Time Coordinated世界协调时间)，又叫做0时区
	fmt.Println(now.In(local2)) // CST(China Standard Time北京时间)，又叫做东8区
	fmt.Println(now.In(local3)) // PDT（美国时区），又叫做西8区
}

// 比较时间大小
func TestTimeCompare(t *testing.T) {
	now := time.Now()
	// 今天10点
	todayClock := time.Date(now.Year(), now.Month(), now.Day(), 10, 0, 0, 0, now.Location())
	// 当前时间在今天10点之后
	b := now.After(todayClock)
	fmt.Println("在今天10点之后", b)
	// 当前时间在今天10点之前
	b = now.Before(todayClock)
	fmt.Println("在今天10点之前", b)
	return
}

func TestTimeAdd(t *testing.T) {
	now := time.Now().Add(1 * time.Hour)
	fmt.Println(now)
}

func TestTimeSubDays(t *testing.T) {
	startTime := time.Now()
	endTime := startTime.Add(48*time.Hour + 1*time.Second)
	// 结束时间取固定时间10点
	endTime = time.Date(endTime.Year(), endTime.Month(), endTime.Day(), 10, 0, 0, 0, time.Local)
	int := timeSubDays(startTime, endTime)
	t.Log(int)
}

// 相隔24小时内为一天
// 相隔非24小时整数，并在结果上加一天
// 相隔24小时，大于1秒都算多一天
func timeSubDays(startTime, endTime time.Time) int {
	subSeconds := endTime.Sub(startTime).Seconds()
	if subSeconds <= 0 {
		return 0
	}
	// 24小时内算一天
	if subSeconds > 0 && subSeconds <= 86400 {
		return 1
	}
	// 对一天的总秒数取余数
	remainder := int(subSeconds) % 86400
	// 除以一天的总秒数
	division := int(subSeconds) / 86400
	// 余数结果是一天的整天数
	if remainder == 0 {
		return division
	} else {
		// 余数结果大于一天并且非整天数
		return division + 1
	}
}

func TestTime(t *testing.T) {
	now := time.Now()
	tenClock := time.Date(now.Year(), now.Month(), now.Day(), 10, 0, 0, 0, now.Location())
	zeroClock := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location())
	if now.After(tenClock) && now.Before(zeroClock) {
		t.Log("true")
	} else {
		t.Log("false")
	}
}

func TestTimeStampNow(t *testing.T) {
	tNow := time.Now().Unix()
	sNow := strconv.FormatInt(tNow, 10)
	t.Log(tNow)
	t.Log(sNow)
}

////////////////////////////////// 用channel实现定时任务 start
var mapChan map[int]<-chan time.Time

func TestTimer(t *testing.T) {
	initTimer()
	resetDuration(1, time.Second*10)
	timer1 := time.NewTimer(time.Second * 1)
	timer2 := time.NewTimer(time.Second * 2)
	timer3 := time.NewTimer(time.Second * 3)

	quitChan := make(chan bool, 1)
	quitChan <- true
	// todo 触发多组定时任务
	for {
		startTimer(quitChan, &Timer{TimerID: 1, Timer: timer1})
		startTimer(quitChan, &Timer{TimerID: 2, Timer: timer2})
		startTimer(quitChan, &Timer{TimerID: 3, Timer: timer3})
	}
	time.Sleep(10 * time.Second)
}

func initTimer() {
	// todo 初始化所有未开始的定时器
	// todo for{}
}

var timerMap map[int]*time.Timer // todo 需要线程安全

type Timer struct {
	TimerID int
	Timer   *time.Timer
}

func startTimer(quit chan bool, timer *Timer) {
	timerMap[timer.TimerID] = timer.Timer
	go func() {
		select {
		case <-timer.Timer.C:
			// todo 检查此任务数据库状态
			// todo 成功后，至掉数据库状态
			delete(timerMap, timer.TimerID)
			break
		case <-quit:
			break
		}
	}()
}

func resetDuration(timerID int, duration time.Duration) {
	if timer, ok := timerMap[timerID]; ok {
		timer.Reset(duration)
	}
}

////////////////////////////////// 用channel实现定时任务 end
