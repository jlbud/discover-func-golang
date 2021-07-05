package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type limiterI interface {
	acquire(ack ...func()) error
	release(done ...func()) error
}

type limiter struct {
	token string
	sem   chan struct{}
	once  sync.Once
}

var m = make(map[string]*limiter)

func addLimiter(token string, n int) {
	m[token] = &limiter{
		sem: make(chan struct{}, n),
	}
}

func (l *limiter) acquire(ack ...func()) error {
	l.once.Do(func() {
		if l.sem == nil {
			if v, ok := m[l.token]; ok {
				l.sem = v.sem
			} else {
				fmt.Printf("not found by token %s\n", l.token)
			}
		}
	})

	select {
	case l.sem <- struct{}{}:
		if len(ack) > 0 {
			ack[0]()
		}
		return nil
	}
}

func (l *limiter) release(done ...func()) error {
	if len(done) > 0 {
		done[0]()
	}
	if l.sem == nil {
		return fmt.Errorf("not found sem by token %s", l.token)
	}
	<-l.sem
	return nil
}

func TestLimiter(t *testing.T) {
	addLimiter("jon", 1)

	li := &limiter{
		token: "aa",
	}
	go func() {
		err := li.acquire()
		if err != nil {
			return
		}
		println("1")
		time.Sleep(1 * time.Second)
		li.release()
	}()

	go func() {
		err := li.acquire()
		if err != nil {
			return
		}
		println("2")
		time.Sleep(1 * time.Second)
		li.release()
	}()

	go func() {
		err := li.acquire()
		if err != nil {
			return
		}
		println("3")
		time.Sleep(1 * time.Second)
		li.release()
	}()

	time.Sleep(4 * time.Second)
	println("this is finished")
}
