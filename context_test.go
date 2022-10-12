package main

import (
	"context"
	"fmt"
	"testing"
	"time"
)

type Context struct {
	ctx    context.Context
	cancel context.CancelFunc
}

func TestContextCancel(t *testing.T) {
	s := Context{}
	s.ctx, s.cancel = context.WithCancel(context.Background())

	go func() {
		resp := te(s.ctx)
		fmt.Println(resp)
	}()
	time.Sleep(2 * time.Second)
	s.cancel()
	time.Sleep(5 * time.Second)
	fmt.Println("test end")
}

func te(ctx context.Context) string {
Done:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("context cancel")
			break Done
		}
	}
	return "eof"
}
