package main

import (
	"bytes"
	"runtime"
	"strconv"
	"testing"
	"time"
)

func TestGroutine1(t *testing.T) {
	t.Log(GetGID())
	go func() {
		t.Logf("go: %d", GetGID())
	}()

	t.Log(GetGID())
	go func() {
		t.Logf("go: %d", GetGID())
	}()
	t.Log(GetGID())
	time.Sleep(2 * time.Second)
}

func GetGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}
