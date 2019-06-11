package main

import (
	"fmt"
	"testing"
	"time"
)

type DataCollectTimer struct {
	durationOneDay   time.Duration
	durationFewHours time.Duration
	execHourTime     int
}

func (dct *DataCollectTimer) startCollectPlan() {
	dct.initTimerParam()

	fmt.Println(dct.durationFewHours)
	fmt.Println(dct.durationOneDay)
	fmt.Println(dct.execHourTime)
	switch dct.isPastTenClock() {
	case true:
		dct.createTimer(dct.durationOneDay)
	case false:
		dct.createTimer(dct.durationFewHours)
	}
}

func (dct *DataCollectTimer) initTimerParam() {
	dct.durationOneDay = time.Hour * 24
	dct.durationFewHours = 0
	dct.execHourTime = 10
}

func (dct *DataCollectTimer) createTimer(duration time.Duration) {
	for {
		now := time.Now()
		// ten o'clock today or tomorrow
		next := now.Add(duration)
		// reset rotate every 24 hours if duration is 0
		if duration == 0 {
			duration = time.Hour * 24
		}
		// ten o'clock
		next = time.Date(next.Year(), next.Month(), next.Day(), dct.execHourTime, 0, 0, 0, next.Location())
		t := time.NewTimer(next.Sub(now))
		fmt.Println("remaining ", next.Sub(now))
		// set timer
		<-t.C
		fmt.Println("do something")
	}
}

func (dct *DataCollectTimer) isPastTenClock() bool {
	now := time.Now()
	// today, ten o'clock
	todayClock := time.Date(now.Year(), now.Month(), now.Day(), dct.execHourTime, 0, 0, 0, now.Location())
	// is past ten o'clock
	return now.After(todayClock)
}

func TestTimePolling(t *testing.T) {
	go func() {
		dct := &DataCollectTimer{}
		dct.startCollectPlan()
	}()

	time.Sleep(10 * time.Hour)
}
