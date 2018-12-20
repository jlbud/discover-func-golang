package main

import (
	"errors"
	"fmt"
	"testing"
)

// 自定义error
func Test_err(t *testing.T) {
	c := &C{
		Err: errors.New("abc"),
	}
	fmt.Println(c.Err.Error())
}

type C struct {
	Err error
}

func (c *C) Error() string {
	err := fmt.Errorf("%s", c.Err.Error())
	return err.Error()
}
