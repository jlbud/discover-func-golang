package main

import (
	"os"
	"os/signal"
	"syscall"
	"testing"
)

// 接收退出信号
func TestSig1(t *testing.T) {
	t.Log("login...")
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
	t.Log("logout...")
}
