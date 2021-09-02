package unisound

import (
	"sync"
	"testing"
	"time"
)

var ch = make(chan int8, 10000)

const MAX = 50

func TestChan(t *testing.T) {

}

func Push(b int8) {
	ch <- b
}

func init() {
	go consumer(ch)
	go consumer(ch)
	go consumer(ch)
}

func consumer(ch chan int8) {
	pr := pusher{
		bufferSli: make([]int8, MAX),
	}
	for {
		select {
		case v, ok := <-ch:
			if ok {
				if pr.isReady() {
					pr.push()
				} else {
					pr.add(v)
				}
			}
		case <-time.After(30 * time.Second):
			pr.push()
		}
	}
}

type pusher struct {
	bufferSli []int8
	mu        sync.RWMutex
}

func (p *pusher) isReady() bool {
	if len(p.bufferSli) >= MAX {
		return true
	}
	return false
}

func (p *pusher) sliSize() int {
	return len(p.bufferSli)
}

func (p *pusher) add(i int8) {
	p.bufferSli = append(p.bufferSli, i)
}

func (p *pusher) push() {
	p.mu.Lock()
	// TODO
	p.bufferSli = p.bufferSli[:0]
	p.mu.Unlock()
}
